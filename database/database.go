package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var d *gorm.DB

func Initialize(dsn string) {
	var err error
	d, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Initialize Database.\n")
}

func GetDB() *gorm.DB {
	return d
}
