package holiday

import (
	"statutory-holidays/pkg/history"
	"strconv"
	"strings"
)

type GetHolidaysParams struct {
	Year   string `form:"year" json:"year"`
	Return string `form:"return" json:"return"`
	ChName string `form:"ch_name" json:"ch_name"`
	EnName string `form:"en_name" json:"en_name"`
}

func (g GetHolidaysParams) validYear() bool {
	year, err := strconv.Atoi(strings.TrimSpace(g.Year))
	if err != nil {
		return false
	}
	return !(year > 2019 || year < 2010)
}

func isChNameExists(value string) bool {
	for _, elem := range history.ChHolidays {
		if strings.Contains(elem, value) {
			return true
		}
	}
	return false
}

func isEnNameExists(value string) bool {
	for _, elem := range history.EnHolidays {
		if strings.Contains(elem, value) {
			return true
		}
	}
	return false
}

func (g GetHolidaysParams) validChName() bool {
	return isChNameExists(g.ChName)

}

func (g GetHolidaysParams) validEnName() bool {
	return isEnNameExists(g.EnName)
}

func (g GetHolidaysParams) CheckQuery() bool {
	return g.validYear() && g.validChName() && g.validEnName()

}
