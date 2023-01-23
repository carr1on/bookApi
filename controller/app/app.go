package app

import (

	"github.com/jinzhu/gorm"                  
	_ "github.com/jinzhu/gorm/dialects/mysql" 

	"log"
)

var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open("mysql", "root:mariadb@/tmp?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Print("failed to connect database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
