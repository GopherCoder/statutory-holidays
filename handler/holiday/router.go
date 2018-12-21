package holiday

import "github.com/gin-gonic/gin"

func RegisterHoliday(r *gin.RouterGroup) {

	r.GET("/holidays", GetHolidaysHandler)
	r.GET("/years/:year", GetHolidaysByHscanHandler)
	r.GET("/holidays/counts/:year", CountYearHandler)
}
