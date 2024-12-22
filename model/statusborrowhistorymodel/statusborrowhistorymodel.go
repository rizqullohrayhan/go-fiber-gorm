package statusborrowhistorymodel

import (
	"github.com/rizqullorayhan/go-fiber-gorm/database"
	"github.com/rizqullorayhan/go-fiber-gorm/entities"
)

func GetAll() ([]entities.StatusBorrowHistory, error) {
	var status []entities.StatusBorrowHistory
	if err := database.DB.Find(&status).Error; err != nil {
		return nil, err
	}
	return status, nil
}

func GetOneByID(statusId uint) (*entities.StatusBorrowHistory, error) {
	var status entities.StatusBorrowHistory
	if err := database.DB.Where("id = ?", statusId).First(&status).Error; err != nil {
		return nil, err
	}
	return &status, nil
}