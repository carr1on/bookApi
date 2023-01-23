package app

import (
	//"gorm.io/driver/sqlite"  //1
	//	"gorm.io/gorm" //1
	"github.com/jinzhu/gorm"                  //2
	_ "github.com/jinzhu/gorm/dialects/mysql" //2
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var (
	db *gorm.DB
)

func Connect() {
	/* //1
	d, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Print("failed to connect database")
	}*/

	//2
	d, err := gorm.Open("mysql", "root:mariadb@/tmp?charset=utf8&parseTime=True&loc=Local")
	// postgreSQL  // _ "github.com/jinzhu/gorm/dialects/postgres"
	//d, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgreSQL")

	//	defer d.Close()
	if err != nil {
		log.Print("failed to connect database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
