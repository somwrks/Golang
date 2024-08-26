package models

import (
	"gorm.io/gorm"
	"github.com/somwrks/Golang/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name   string `gorm:"" json:"name"`
	Author string `json:"author"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}