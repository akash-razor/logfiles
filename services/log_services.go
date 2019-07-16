package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sync"
	"../models"
	"../database"
)

func CreateLog(c *gin.Context){

	var logs []models.LogModel
	c.BindJSON(&logs)
	var wg sync.WaitGroup
	wg.Add(len(logs))
	for _, item := range logs{
		go func(item1 models.LogModel){
			fmt.Println(item1)
			database.Db.Save(&item1)
			wg.Done()
		}(item)

	}
	wg.Wait()
	fmt.Println("done")
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "done"})

}

func ReadUserLog(c *gin.Context){
	var logs []models.LogModel
	descQuery := fmt.Sprintf("%%%v%%", c.Query("description"))
	database.Db.Where("description LIKE ?", descQuery).Find(&logs)
	if len(logs) <= 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message":"No log found!"})
		return
	}

	for i:=0 ; i<len(logs); i++{
		logs[i].Name = "xxxxxxx"
		logs[i].Email = "xxxxxxx"
		logs[i].Phone = 000000000

	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data":logs})
}

func ReadAdminLog(c *gin.Context){
	var logs []models.LogModel
	nameQuery := fmt.Sprintf("%%%v%%", c.Query("name"))
	emailQuery := fmt.Sprintf("%%%v%%", c.Query("email"))
	descQuery := fmt.Sprintf("%%%v%%", c.Query("description"))
	database.Db.Where("name LIKE ?", nameQuery).Or("email LIKE ?", emailQuery).Or("phone = ?", c.Query("phone")).Or("description = ?", descQuery).Find(&logs)
	if len(logs) <= 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message":"No todo found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data":logs})
}

func FileCreateLog(c *gin.Context){
	plan, _ := ioutil.ReadFile(c.Query("filePath"))
	var logs []models.LogModel
	json.Unmarshal(plan, &logs)
	var wg sync.WaitGroup
	wg.Add(len(logs))
	for _, item := range logs{
		go func(item1 models.LogModel){
			fmt.Println(item1)
			database.Db.Save(&item1)
			wg.Done()
		}(item)

	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data":logs})
}
