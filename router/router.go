package router

import (
	"net/http"
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
	{
		g.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, fetchAllPath(g))
		})
	}
	return g
}

func fetchAllPath(g *gin.Engine) []string {
	routers := g.Routes()
	var paths []string
	for _, router := range routers {
		paths = append(paths, router.Path)
	}
	return paths
}
