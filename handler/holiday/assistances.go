package holiday

import (
	"math/rand"
	"statutory-holidays/pkg/history"
	"statutory-holidays/pkg/initial"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

func FetchKeyByYearReturnAll(year string) []string {
	keys, _ := redis.Strings(initial.RedisConn.Do("HKEYS", initial.HistoryKey))
	var yearKeys []string
	for _, key := range keys {
		if strings.Contains(key, year) {
			yearKeys = append(yearKeys, key)
		}
	}
	return yearKeys

}

func FetchKeyByYearReturnOne(year string) string {
	var yearKeys []string
	yearKeys = FetchKeyByYearReturnAll(year)
	rand.Seed(time.Now().UnixNano())
	return yearKeys[rand.Intn(len(yearKeys))]

}

func fetchChNameIndex(value string) int {
	for index, elem := range history.ChHolidays {
		if elem == value {
			return index
		}
	}
	return -1
}

func fetchEnNameIndex(value string) int {
	for index, elem := range history.EnHolidays {
		if elem == value {
			return index
		}
	}
	return -1
}

func splitYearKey(value string) (year int, index int) {
	stringList := strings.Split(value, ":")
	year, _ = strconv.Atoi(stringList[0])
	index, _ = strconv.Atoi(stringList[1])
	return
}

func FetchKeyByChNameReturnOne(year string, chName string) string {
	var yearKeys []string
	yearKeys = FetchKeyByYearReturnAll(year)
	index := fetchChNameIndex(chName)
	for _, year := range yearKeys {
		_, i := splitYearKey(year)
		if i == index {
			return year
		}
	}
	return "-1"

}

func FetchKeyByEnNameReturnOne(year string, enName string) string {
	var yearKeys []string
	yearKeys = FetchKeyByYearReturnAll(year)
	index := fetchEnNameIndex(enName)
	for _, year := range yearKeys {
		_, i := splitYearKey(year)
		if i == index {
			return year
		}
	}
	return "-1"
}

func Count(value string) int {
	stringList := strings.Split(value, "~")
	left, _ := time.Parse("2006/01/02", stringList[0])
	right, _ := time.Parse("2006/01/02", stringList[1])
	return int(right.Sub(left).Hours()/24) + 1

}

func KeyHandler(value string) []string {
	/*
		index: 0  : ch_name
		index: 1  : en_name
	*/
	stringList := strings.Split(value, ":")
	index, _ := strconv.Atoi(stringList[1])
	var result []string
	result = append(result, history.ChHolidays[index])
	result = append(result, history.EnHolidays[index])
	return result

}
