package holiday

import "github.com/gin-gonic/gin"

func RegisterHoliday(r *gin.RouterGroup) {

	r.GET("/holidays", GetHolidaysHandler)
	r.GET("/holidays/counts")
}
