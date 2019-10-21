package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var GromDb *gorm.DB

func init() {
	var err error
	GromDb, err = gorm.Open("mysql", "root:root@/go_study?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("db连接情况:", GromDb)
}
