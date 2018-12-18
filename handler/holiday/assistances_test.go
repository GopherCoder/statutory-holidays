package holiday

import (
	"fmt"
	"statutory-holidays/pkg/initial"
	"testing"
)

func TestFetchKeyByYearReturnAll(t *testing.T) {
	initial.Start()
	tt := []struct {
		year string
	}{
		{
			year: "2019",
		},
		{
			year: "2018",
		},
		{
			year: "2017",
		},
		{
			year: "2016",
		},
	}
	for _, t := range tt {
		fmt.Println(FetchKeyByYearReturnAll(t.year))
	}
}

func TestFetchKeyByYearReturnOne(t *testing.T) {
	initial.Start()
	tt := []struct {
		year string
	}{
		{
			year: "2019",
		},
		{
			year: "2018",
		},
	}
	for _, t := range tt {
		fmt.Println(FetchKeyByYearReturnOne(t.year))
	}
}

func TestFetchChNameIndex(t *testing.T) {
	tt := []struct {
		name string
	}{
		{
			name: "元旦",
		},
		{
			name: "中秋节",
		},
	}
	for _, t := range tt {
		fmt.Println(fetchChNameIndex(t.name))
	}
}

func TestFetchEnNameIndex(t *testing.T) {
	tt := []struct {
		name string
	}{
		{
			name: "Spring Festival",
		},
		{
			name: "Labour Day",
		},
	}
	for _, t := range tt {
		fmt.Println(fetchEnNameIndex(t.name))
	}
}

func TestSplitYearKey(t *testing.T) {
	tt := []struct {
		key string
	}{
		{
			key: "2019:0",
		},
		{
			key: "2018:12",
		},
	}
	for _, t := range tt {
		fmt.Println(splitYearKey(t.key))
	}
}

func TestFetchKeyByChNameReturnOne(t *testing.T) {
	initial.Start()
	tt := []struct {
		year   string
		chName string
	}{
		{
			year:   "2019",
			chName: "元旦",
		},
		{
			year:   "2018",
			chName: "中秋节",
		},
	}
	for _, t := range tt {
		fmt.Println(FetchKeyByChNameReturnOne(t.year, t.chName))
	}
}

func TestFetchKeyByEnnameReturnOne(t *testing.T) {
	initial.Start()
	tt := []struct {
		year string
		name string
	}{
		{
			year: "2019",
			name: "New Year\\'s Day",
		},
		{
			year: "2018",
			name: "National Day",
		},
	}
	for _, t := range tt {
		fmt.Println(FetchKeyByEnNameReturnOne(t.year, t.name))
	}
}

func TestCount(t *testing.T) {
	tt := []struct {
		value string
	}{
		{
			value: "2018/12/30~2019/01/01",
		},
		{
			value: "2019/02/04~2019/02/10",
		},
	}
	for _, t := range tt {
		fmt.Println(Count(t.value))
	}
}
