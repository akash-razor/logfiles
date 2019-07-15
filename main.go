package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"./database"
	"./routes"
)



func init(){
	database.InitDb()
}



func main(){

	routes.InitRoutes()

}
