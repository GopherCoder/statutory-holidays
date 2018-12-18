package router

import (
	"statutory-holidays/handler/holiday"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) InitRouter(g *gin.Engine, handler ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())
	g.Use(handler...)

	holidays := g.Group("/v1/api/holiday")
	{
		holiday.RegisterHoliday(holidays)
	}
	return g
}
