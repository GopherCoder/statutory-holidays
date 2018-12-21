package holiday

import (
	"fmt"
	"net/http"
	"statutory-holidays/models"
	"statutory-holidays/pkg/initial"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

/*
HSCAN history_holidays_map 0 MATCH *2019*
*/
func GetHolidaysHandler(c *gin.Context) {
	var params GetHolidaysParams
	if err := c.ShouldBindQuery(&params); err != nil {
		return
	}
	if !params.CheckQuery() {
		c.JSON(http.StatusBadGateway, gin.H{
			"data": "check year in (2010~2019)",
		})
		return
	}
	if params.Year != "" {
		var result models.Holidays
		yearKeys := FetchKeyByYearReturnAll(params.Year)
		for _, yearKey := range yearKeys {
			_, index := splitYearKey(yearKey)
			value, _ := redis.String(initial.RedisConn.Do("HGET", initial.HistoryKey, yearKey))
			var one models.Holiday
			one.ChName, _ = redis.String(initial.RedisConn.Do("LINDEX", initial.ChNameKey, index))
			one.EnName, _ = redis.String(initial.RedisConn.Do("LINDEX", initial.EnNameKey, index))
			one.Date = value
			one.Count = Count(value)
			result = append(result, one)
		}
		c.JSON(http.StatusOK, result)
		return
	}

}

// HSCAN
func GetHolidaysByHscanHandler(c *gin.Context) {
	var params string
	params = c.Param("year")

	values, err := redis.Values(initial.RedisConn.Do(
		"HSCAN", initial.HistoryKey, 0, "MATCH", fmt.Sprintf("*%s*", params)))

	if err != nil {
		return
	}
	var KeyValues = make(map[string]string)
	KeyValues, _ = redis.StringMap(values[1], nil)

	var result models.Holidays

	for key, value := range KeyValues {
		var one models.Holiday
		tmp := KeyHandler(key)
		one.ChName = tmp[0]
		one.EnName = tmp[1]
		one.Date = value
		one.Count = Count(value)
		result = append(result, one)

	}
	c.JSON(http.StatusOK, result)

}

func CountYearHandler(c *gin.Context) {
	var params string
	params = c.Param("year")

	values, _ := redis.Values(
		initial.RedisConn.Do(
			"HSCAN", initial.HistoryKey, 0, "MATCH", fmt.Sprintf("*%s*", params)))

	var KeyValues = make(map[string]string)
	KeyValues, _ = redis.StringMap(values[1], nil)
	var count int
	for _, value := range KeyValues {
		count += Count(value)
	}
	c.JSON(http.StatusOK,
		gin.H{
			"year":  params,
			"count": count,
		})
}
