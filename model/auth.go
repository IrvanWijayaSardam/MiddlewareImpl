package model

import (
	"gorm.io/gorm"
)

type Login struct {
	Email    string `gorm:"type:varchar(255)"`
	Password string
}

type AuthModel struct {
	db *gorm.DB
}

func (am *AuthModel) Init(db *gorm.DB) {
	am.db = db
}

func (am *AuthModel) Login(dataLogin Login) (bool, error) {
	var user Users

	result := am.db.Where("email = ?", dataLogin.Email).First(&user)
	if result.Error != nil {
		return false, result.Error
	}

	if dataLogin.Password != user.Password {
		return false, nil
	}

	return true, nil
}
