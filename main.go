package main

import (
	"fmt"
	"statutory-holidays/pkg/initial"
	"statutory-holidays/router"

	"github.com/gin-gonic/gin"
)

func main() {
	initial.Start()
	fmt.Println("Hello world")
	var routers router.Router
	g := gin.Default()
	routers.InitRouter(g)
	g.Run(":8080")

}
