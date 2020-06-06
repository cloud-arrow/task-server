package model

import (
	"fmt"
	"github.com/cloud-arrow/task-server/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func init() {
	var err error
	var user = config.Config.Database.User
	var password = config.Config.Database.Password
	var dbname = config.Config.Database.DBname
	url := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbname)
	db, err = gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
}
