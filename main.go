package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"net/http"
	"sync"

	//"sync"


	//"encoding/json"
)

var db *gorm.DB


func init(){
	var err error
	db, err = gorm.Open("mysql", "root:12345678@/logfiles?charset=utf8&parseTime=True&loc=Local")
	if err!= nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&logModel{})
}

type(
	logModel struct{
		gorm.Model
		Name string `json:"name"`
		Email string `json:"email"`
		Phone int `json:"phone"`
		Description string `json:"description"`
	}

	obfucatedLog struct{
		Name string `json:"name"`
		Email string `json:"email"`
		Phone int `json:"phone"`
		Description string `json:"description"`
		FileId string `json:"file_id"`
	}
)


func main(){

	router := gin.Default()

	v1 := router.Group("/api/v1/logs")
	{
		v1.POST("/", createLog)
		v1.GET("/user/", readUserLog)
		v1.GET("/admin/", readAdminLog)
		v1.GET("/file/", fileCreateLog)
	}

	router.Run()

}

func createLog(c *gin.Context){

	//fmt.Println(c)
	var logs []logModel
	c.BindJSON(&logs)
	var wg sync.WaitGroup
	wg.Add(len(logs))
	for _, item := range logs{
		//fmt.Println(item)
		//db.Save(&item)
		//fmt.Println(item)
		go func(item1 logModel){
			fmt.Println(item1)
			db.Save(&item1)
			wg.Done()
		}(item)

	}
	wg.Wait()
	fmt.Println("done")
	//db.Create(logs)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "done"})

}

func readUserLog(c *gin.Context){
	var logs []logModel
	//var _logs []obfucatedLog
	db.Where("description LIKE ?", c.Query("description")).Find(&logs)
	if len(logs) <= 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message":"No todo found!"})
		return
	}

	for i:=0 ; i<len(logs); i++{
		logs[i].Name = "xxxxxxx"
		logs[i].Email = "xxxxxxx"
		logs[i].Phone = 000000000
		//fmt.Println(item)
		//Name := "xxxxxxx"
		//Email := "xxxxxxx"
		//Phone := 00000000

	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data":logs})
}

func readAdminLog(c *gin.Context){
	var logs []logModel
	db.Where("name LIKE ?", c.Query("name")).Or("email LIKE ?", c.Query("email")).Or("phone = ?", c.Query("phone")).Or("description = ?", c.Query("description")).Find(&logs)
	//db.Find(&logs)
	if len(logs) <= 0{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message":"No todo found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data":logs})
}

func fileCreateLog(c *gin.Context){
	plan, _ := ioutil.ReadFile(c.Query("filePath"))
	//fmt.Println("query",c.Query("sdhsdv"))
	if c.Query("dfgf")==""{
		fmt.Println("here")
	}
	var logs []logModel
	json.Unmarshal(plan, &logs)
	var wg sync.WaitGroup
	wg.Add(len(logs))
	for _, item := range logs{
		//fmt.Println(item)
		//db.Save(&item)
		//fmt.Println(item)
		go func(item1 logModel){
			fmt.Println(item1)
			db.Save(&item1)
			wg.Done()
		}(item)

	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data":logs})
}