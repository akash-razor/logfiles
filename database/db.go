package database

import (
	"github.com/jinzhu/gorm"
	"../models"
)

var Db *gorm.DB

func InitDb(){
	var err error
	Db, err = gorm.Open("mysql", "root:12345678@/logfiles?charset=utf8&parseTime=True&loc=Local")
	if err!= nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&models.LogModel{})
}
