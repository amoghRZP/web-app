package main

import (
	"fmt"
	"github.com/amoghRZP/web-app/dbconfig"
	"github.com/amoghRZP/web-app/models"
	"github.com/amoghRZP/web-app/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var err error

func main() {
	//Commenting sqlite to use mysql
	//db, err = gorm.Open("sqlite3", "./gorm.db")
	dbconfig.DB, err = gorm.Open("mysql", "root:amogh123@tcp(127.0.0.1:3306)/Register?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}

	defer dbconfig.DB.Close()
	dbconfig.DB.AutoMigrate(&models.Notes{})

	r := routers.GetRouter()

	r.Run(":8080")
}
