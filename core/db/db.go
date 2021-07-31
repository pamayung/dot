package db

import (
	"dot/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func configDB() *gorm.DB {
	USER := "root"
	PASS := "12345"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "dot"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err = gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func ConnectDB() {
	db := configDB()
	db.AutoMigrate(&model.Account{})
	db.AutoMigrate(&model.Member{})
}

func DbManager() *gorm.DB {
	return db
}
