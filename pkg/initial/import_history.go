package initial

import (
	"errors"
	"fmt"
	"statutory-holidays/pkg/history"

	"github.com/gomodule/redigo/redis"
)

/*
chName:
	list:
		- 1:
		- 2:

enName:
	list:
		- 1:
		- 2:
history:
	map:
		key:value 2019:1 2019.01.01~2019.01.03
		key:value 2019:2 2019.01.01~2019.01.03
		key:value 2019:0 2019.01.01~2019.01.03

*/

const (
	ChNameKey  string = "ch_name_list"
	EnNameKey  string = "en_name_list"
	HistoryKey string = "history_holidays_map"
)

var RedisConn redis.Conn

func initDial() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic("connect redis fail")
	}

	RedisConn = c
	return RedisConn
}

func isKeyExists(value string) (bool, error) {
	return redis.Bool(RedisConn.Do("EXISTS", value))
}

func chNameImport() (bool, error) {
	if ok, _ := isKeyExists(ChNameKey); ok {
		return false, errors.New("ch_name_list not exists")
	}
	chNameCollections := history.ChHolidays
	for _, collection := range chNameCollections {
		if ok, _ := redis.Bool(RedisConn.Do("RPUSH", ChNameKey, collection)); !ok {
			panic("RPUSH ch_name_list fail")
		}
	}
	return true, nil
}

func enNameImport() (bool, error) {
	if ok, _ := isKeyExists(EnNameKey); ok {
		return false, errors.New("en_name_list not exists")
	}
	enNameCollections := history.EnHolidays
	for _, collection := range enNameCollections {
		if ok, _ := redis.Bool(RedisConn.Do("RPUSH", EnNameKey, collection)); !ok {
			panic("RPUSH en_name_list fail")
		}
	}
	return true, nil

}

func historyImport() (bool, error) {
	if ok, _ := isKeyExists(HistoryKey); ok {
		return false, errors.New(HistoryKey + " not exists")
	}
	collections := history.FetchCollectionYearHistory().Data
	for yearIndex, yearHolidays := range collections {
		for holidayIndex, holiday := range yearHolidays {
			yearField := 2019 - yearIndex
			yearValue := fmt.Sprintf("%s~%s", holiday.Start, holiday.End)
			fmt.Println(HistoryKey, fmt.Sprintf("%d:%d", yearField, holidayIndex), yearValue)
			if ok, _ := redis.Bool(RedisConn.Do("HSET", HistoryKey, fmt.Sprintf("%d:%d", yearField, holidayIndex), yearValue)); !ok {
				panic("HSET fail")
			}
		}
	}
	return true, nil
}

func Start() {
	initDial()
	if _, err := chNameImport(); err != nil {
		return
	}

	if _, err := enNameImport(); err != nil {
		return
	}
	if _, err := historyImport(); err != nil {
		return
	}

}
