package controllers

import (
	"github.com/gin-gonic/gin"
	"../services"
)

func CreateLog(c *gin.Context){
	services.CreateLog(c)
}

func ReadUserLog(c *gin.Context){
	services.ReadUserLog(c)
}

func ReadAdminLog(c *gin.Context){
	services.ReadAdminLog(c)
}

func FileCreateLog(c *gin.Context){
	services.FileCreateLog(c)
}
