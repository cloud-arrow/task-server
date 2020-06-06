package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Nickname string
	Phone string
}

func AddUser(user *User){
	db.Create(user)
}
