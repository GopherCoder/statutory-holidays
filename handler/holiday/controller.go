package holiday

import (
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
func GetHolidaysHscanHandler(c *gin.Context) {}

func CountNumberHandler(c *gin.Context) {}
