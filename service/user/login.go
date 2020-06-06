package user

import (
	"errors"
	"github.com/cloud-arrow/task-server/model"
)

type UserDto struct {
	Phone    string
	Password string
}

func Login(userDto *UserDto) error {
	if userDto.Phone != "18065048384" {
		return errors.New("手机或密码错误")
	}
	userPo := &model.User{Nickname: "fangqy", Phone: userDto.Phone}
	model.AddUser(userPo)
	return nil

}
