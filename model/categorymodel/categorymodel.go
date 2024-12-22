package categorymodel

import (
	"github.com/rizqullorayhan/go-fiber-gorm/database"
	"github.com/rizqullorayhan/go-fiber-gorm/dto"
	"github.com/rizqullorayhan/go-fiber-gorm/entities"
	"gorm.io/gorm/clause"
	// "gorm.io/gorm"
)

func GetAll() ([]dto.GetCategory, error) {
	var categories []entities.Category
	if err := database.DB.Preload("Books").Find(&categories).Error; err != nil {
		return nil, err
	}

	var categoryDTOs []dto.GetCategory
	for _, category := range categories {
		categoryDTO := dto.GetCategory{
			ID:   category.ID,
			Name: category.Name,
			Books: category.Books,
		}
		categoryDTOs = append(categoryDTOs, categoryDTO)
	}
	return categoryDTOs, nil
}

func GetOneByID(categoryId uint) (*entities.Category, error) {
	var category entities.Category

	if err := database.DB.Preload("Books").Where("id = ?", categoryId).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func GetManyByID(categoriesId []uint) (*[]entities.Category, error) {
	var categories []entities.Category

	if err := database.DB.Where("id IN ?", categoriesId).Find(&categories).Error; err != nil {
		return nil, err
	}
	return &categories, nil
}

func Create(category *dto.CreateCategory) error {
	newCategory := entities.Category{
		Name: category.Name,
	}

	return database.DB.Create(&newCategory).Error
}

func Update(category *entities.Category) error {
	return database.DB.Save(category).Error
}

func Delete(categoryId uint) error {
	return database.DB.Select(clause.Associations).Delete(&entities.Category{ID:categoryId}).Error
}
