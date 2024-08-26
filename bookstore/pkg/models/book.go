package models

import (
	"github.com/jinzhu/gorm"
	"github.com/somwrks/Golang/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

