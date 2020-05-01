package service

import (
	"errors"

	"authen/model"

	log "github.com/alecthomas/log4go"
	"github.com/jinzhu/gorm"
)

// UserRegisterSvc 注册服务接口json
type UserRegisterSvc struct {
	LoginName string `json:"username"`
	Passwd    string `json:"passwd"`
}

// Register 用户注册实现
func (service *UserRegisterSvc) Register() (model.User, error) {

	user := model.User{
		LoginName: service.LoginName,
	}

	err := model.DB.Transaction(func(tx *gorm.DB) error {
		count := 0
		if tx.Model(&model.User{}).Where("login_name = ?", service.LoginName).Count(&count); count == 0 {
			if err := user.SetPasswd(service.Passwd); err != nil {
				log.Error(err)
				return errors.New("set password failed")
			}
			if err := tx.Create(&user).Error; err != nil {
				log.Error(err)
				return errors.New("create db record failed")
			}
			return nil
		}
		log.Error("login name already exists")
		return errors.New("login name already exists")
	})

	return user, err
}

// UserLoginSvc 登录服务接口json
type UserLoginSvc struct {
	LoginName string `json:"username"`
	Passwd    string `json:"passwd"`
}

// Login 用户登录实现
func (service *UserLoginSvc) Login() (model.User, error) {
	var user model.User
	if err := model.DB.Where("login_name = ?", service.LoginName).First(&user).Error; err == nil {
		if isValid := user.CheckPasswd(service.Passwd); isValid {
			log.Info("user %s login", user.LoginName)
			return user, nil
		}
		return user, errors.New("password incorrect")
	} else {
		log.Error(err)
		return user, errors.New("login name not found")
	}
}
