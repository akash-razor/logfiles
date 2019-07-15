package routes

import (
	"github.com/gin-gonic/gin"
	"../controllers"
)

func InitRoutes(){
	router := gin.Default()

	v1 := router.Group("/api/v1/logs")
	{
		v1.POST("/", controllers.CreateLog)
		v1.GET("/user/", controllers.ReadUserLog)
		v1.GET("/admin/", controllers.ReadAdminLog)
		v1.GET("/file/", controllers.FileCreateLog)
	}

	router.Run()
}
