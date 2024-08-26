package config

import(
	"github.com/jinzhu/gorm"
	"gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var (
	db *gorm.DB
)

func Connect(){
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}