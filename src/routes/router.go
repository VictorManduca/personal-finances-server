package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	controllers "github.com/victormanduca/personal-finances/src/controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/", controllers.Calc)

	return router
}
