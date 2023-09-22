package model

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Books struct {
	ID          string `gorm:"primaryKey;type:varchar(100)"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Date        string `gorm:"type:varchar(100)"`
	Publisher   string `gorm:"type:varchar(255)"`
}

func (u *Books) GenerateID() {
	if u.ID == "" {
		u.ID = uuid.NewString()
		return
	}
	fmt.Println("ID sudah ada")
}

type BooksModel struct {
	db *gorm.DB
}

func (bm *BooksModel) Init(db *gorm.DB) {
	bm.db = db
}

func (bm *BooksModel) Create(newBook Books) *Books {
	newBook.GenerateID()
	if err := bm.db.Create(&newBook).Error; err != nil {
		return nil
	}
	return &newBook
}

func (bm *BooksModel) Get(id string) *Books {
	var book Books
	if err := bm.db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil
	}
	return &book
}

func (bm *BooksModel) GetAll() []Books {
	var books []Books
	if err := bm.db.Find(&books).Error; err != nil {
		return nil
	}
	return books
}

func (bm *BooksModel) Delete(id int) {
	var deletdData = Books{}
	deletdData.ID = strconv.Itoa(id)
	if err := bm.db.Delete(&deletdData).Error; err != nil {
		logrus.Error("Model : Delete error, ", err.Error())
	}
}

func (bm *BooksModel) UpdateData(updatedData Books) bool {
	var qry = bm.db.Save(&updatedData)
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
