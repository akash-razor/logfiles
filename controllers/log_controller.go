package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sync"
	"../models"
	"../database"
	"../services"
)

func CreateLog(c *gin.Context){

	services.CreateLog(c)

	//fmt.Println(c)
	//var logs []models.LogModel
	//c.BindJSON(&logs)
	//var wg sync.WaitGroup
	//wg.Add(len(logs))
	//for _, item := range logs{
	//	//fmt.Println(item)
	//	//db.Save(&item)
	//	//fmt.Println(item)
	//	go func(item1 models.LogModel){
	//		fmt.Println(item1)
	//		database.Db.Save(&item1)
	//		wg.Done()
	//	}(item)
	//
	//}
	//wg.Wait()
	//fmt.Println("done")
	////db.Create(logs)
	//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "done"})

}

func ReadUserLog(c *gin.Context){
	services.ReadUserLog(c)
	//var logs []models.LogModel
	////var _logs []obfucatedLog
	//database.Db.Where("description LIKE ?", c.Query("description")).Find(&logs)
	//if len(logs) <= 0{
	//	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message":"No todo found!"})
	//	return
	//}
	//
	//for i:=0 ; i<len(logs); i++{
	//	logs[i].Name = "xxxxxxx"
	//	logs[i].Email = "xxxxxxx"
	//	logs[i].Phone = 000000000
	//	//fmt.Println(item)
	//	//Name := "xxxxxxx"
	//	//Email := "xxxxxxx"
	//	//Phone := 00000000
	//
	//}
	//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data":logs})
}

func ReadAdminLog(c *gin.Context){
	var logs []models.LogModel
	database.Db.Where("name LIKE ?", c.Query("name")).Or("email LIKE ?", c.Query("email")).Or("phone = ?", c.Query("phone")).Or("description = ?", c.Query("description")).Find(&logs)
	//db.Find(&logs)
	if len(logs) <= 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message":"No todo found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data":logs})
}

func FileCreateLog(c *gin.Context){
	plan, _ := ioutil.ReadFile(c.Query("filePath"))
	//fmt.Println("query",c.Query("sdhsdv"))
	if c.Query("dfgf")==""{
		fmt.Println("here")
	}
	var logs []models.LogModel
	json.Unmarshal(plan, &logs)
	var wg sync.WaitGroup
	wg.Add(len(logs))
	for _, item := range logs{
		//fmt.Println(item)
		//db.Save(&item)
		//fmt.Println(item)
		go func(item1 models.LogModel){
			fmt.Println(item1)
			database.Db.Save(&item1)
			wg.Done()
		}(item)

	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data":logs})
}
