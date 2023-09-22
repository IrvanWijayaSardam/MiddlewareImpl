package model

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Users struct {
	ID       string `gorm:"primaryKey;type:varchar(100)"`
	Name     string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255)"`
	Password string
}

func (u *Users) GenerateID() {
	if u.ID == "" {
		u.ID = uuid.NewString()
		return
	}
	fmt.Println("ID sudah ada")
}

type UsersModel struct {
	db *gorm.DB
}

func (um *UsersModel) Init(db *gorm.DB) {
	um.db = db
}

func (um *UsersModel) Create(newUser Users) *Users {
	newUser.GenerateID()
	if err := um.db.Create(&newUser).Error; err != nil {
		logrus.Error("Model : Insert data error, ", err.Error())
		return nil
	}

	return &newUser
}

func (um *UsersModel) Get(id string) Users {
	var user Users
	if err := um.db.Where("id = ?", id).First(&user).Error; err != nil {
		logrus.Error("Model: Error fetching user by ID, ", err.Error())
	}

	return user
}

func (um *UsersModel) GetAll() []Users {
	var listUser = []Users{}
	if err := um.db.Find(&listUser).Error; err != nil {
		logrus.Error("Model : Insert data error, ", err.Error())
		return nil
	}

	return listUser
}

func (um *UsersModel) UpdateData(updatedData Users) bool {
	var qry = um.db.Save(&updatedData)
	if err := qry.Error; err != nil {
		logrus.Error("Model : Update error, ", err.Error())
		return false
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		logrus.Error("Model : Update error, ", "no data affected")
		return false
	}

	return true
}

func (um *UsersModel) Delete(id int) {
	var deletdData = Users{}
	deletdData.ID = strconv.Itoa(id)
	if err := um.db.Delete(&deletdData).Error; err != nil {
		logrus.Error("Model : Delete error, ", err.Error())
	}
}
