package bookmodel

import (
	"github.com/rizqullorayhan/go-fiber-gorm/database"
	"github.com/rizqullorayhan/go-fiber-gorm/dto"
	"github.com/rizqullorayhan/go-fiber-gorm/entities"
	"github.com/rizqullorayhan/go-fiber-gorm/model/categorymodel"
	"github.com/rizqullorayhan/go-fiber-gorm/utils"
	"gorm.io/gorm/clause"
	// "gorm.io/gorm"
)

func GetAll() ([]dto.GetBook, error) {
	var books []entities.Book
	if err := database.DB.Preload("Categories").Order("title").Find(&books).Error; err != nil {
		return nil, err
	}

	var bookDTOs []dto.GetBook
	for _, book := range books {
		bookDTO := dto.GetBook{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Cover:       book.Cover,
			Categories:  book.Categories,
			IsAvailable: book.IsAvailable,
		}
		bookDTOs = append(bookDTOs, bookDTO)
	}
	return bookDTOs, nil
}

func GetOneByID(bookId uint) (*entities.Book, error) {
	var book entities.Book

	if err := database.DB.Preload("Categories").Where("id = ?", bookId).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func GetManyByID(booksId []uint) (*[]entities.Book, error) {
	var books []entities.Book

	if err := database.DB.Where("id IN ?", booksId).Order("title").Find(&books).Error; err != nil {
		return nil, err
	}
	return &books, nil
}

func Create(book *dto.CreateBook) error {
	categoryIDs := utils.ConvertToUintArray(book.Categories)
	categories, err := categorymodel.GetManyByID(categoryIDs)
	if err != nil {
		return err
	}
	newBook := entities.Book{
		Title:       book.Title,
		Author:      book.Author,
		Cover:       book.Cover,
		IsAvailable: book.IsAvailable,
		Categories:  *categories,
	}

	return database.DB.Create(&newBook).Error
}

func UpdateWithCategory(book *entities.Book, categoriesID []int) error {
	categoryIDs := utils.ConvertToUintArray(categoriesID)
	categories, err := categorymodel.GetManyByID(categoryIDs)
	if err != nil {
		return err
	}
	if err := database.DB.Model(&book).Association("Categories").Clear(); err != nil {
		return err
	}
	book.Categories = *categories
	return database.DB.Save(book).Error
}

func Update(book *entities.Book) error {
	return database.DB.Save(book).Error
}

func Delete(bookId uint) error {
	return database.DB.Select(clause.Associations).Delete(&entities.Book{ID: bookId}).Error
}
