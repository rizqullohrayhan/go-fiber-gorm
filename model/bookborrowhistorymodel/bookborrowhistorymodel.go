package bookborrowhistorymodel

import (
	"errors"
	"time"

	"github.com/rizqullorayhan/go-fiber-gorm/database"
	"github.com/rizqullorayhan/go-fiber-gorm/dto"
	"github.com/rizqullorayhan/go-fiber-gorm/entities"
	"github.com/rizqullorayhan/go-fiber-gorm/model/bookmodel"
	"github.com/rizqullorayhan/go-fiber-gorm/model/statusborrowhistorymodel"
	"github.com/rizqullorayhan/go-fiber-gorm/model/usermodel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAll() ([]entities.BookBorrowHistory, error) {
	var history []entities.BookBorrowHistory
	if err := database.DB.Preload(clause.Associations).Find(&history).Error; err != nil {
		return nil, err
	}

	return history, nil
}

func GetOneByID(historyId uint) (*entities.BookBorrowHistory, error) {
	var history entities.BookBorrowHistory
	err := database.DB.Preload(clause.Associations).Where("id = ?", historyId).First(&history).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // Data tidak ditemukan, kembalikan nil
	}

	if err != nil {
		return nil, err
	}

	return &history, nil
}

func GetByBook(bookId uint) ([]entities.BookBorrowHistory, error) {
	var history []entities.BookBorrowHistory
	if err := database.DB.Preload(clause.Associations).Where("book_id = ?", bookId).Find(&history).Error; err != nil {
		return nil, err
	}

	return history, nil
}

func GetByUser(userId uint) ([]entities.BookBorrowHistory, error) {
	var history []entities.BookBorrowHistory
	if err := database.DB.Preload(clause.Associations).Where("user_id = ?", userId).Find(&history).Error; err != nil {
		return nil, err
	}

	return history, nil
}

func GetByUserAndBook(userId uint, bookId uint) ([]entities.BookBorrowHistory, error) {
	var history []entities.BookBorrowHistory
	if err := database.DB.Preload(clause.Associations).Where("user_id = ? AND book_id = ?", userId, bookId).Find(&history).Error; err != nil {
		return nil, err
	}

	return history, nil
}

func GetUserLatestBorrowBook(userId uint) ([]entities.Book, error) {
	var bookIDs []uint
	if err := database.DB.Model(&entities.BookBorrowHistory{}).Where("user_id = ?", userId).Group("book_id").Order("borrowed_at DESC").Pluck("book_id", &bookIDs).Error; err != nil {
		return nil, err
	}

	books, err := bookmodel.GetManyByID(bookIDs)
	if err != nil {
		return nil, err
	}

	return *books, nil
}

func Create(history *dto.CreateBookBorrowHistory) error {
	book, err := bookmodel.GetOneByID(history.BookID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("buku tidak ditemukan")
	}
	if !book.IsAvailable {
		return errors.New("buku tidak tersedia")
	}
	if _, err := usermodel.GetOneByID(history.UserID); errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user tidak ditemukan")
	}
	if _, err := statusborrowhistorymodel.GetOneByID(history.StatusID); errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("status tidak ditemukan")
	}

	var statusID uint
	if history.ReturnedAt != nil {
		statusID = 2
	} else {
		statusID = history.StatusID
	}

	// Jika BorrowedAt tidak diisi, tetapkan nilai waktu sekarang
	if history.BorrowedAt.IsZero() {
		history.BorrowedAt = time.Now()
	}


	historyNew := entities.BookBorrowHistory{
		BookID: history.BookID,
		UserID: history.UserID,
		StatusID: statusID,
		BorrowedAt: history.BorrowedAt,
		ReturnedAt: history.ReturnedAt,
		Keterangan: history.Keterangan,
	}

	if err := database.DB.Create(&historyNew).Error; err != nil {
		return err
	}

	if statusID == 1 {
		book.IsAvailable = false
	} else {
		book.IsAvailable = true
	}

	return bookmodel.Update(book)
}

func Update(historyID uint, historyNew *dto.UpdateBookBorrowHistory) error {
	history, err := GetOneByID(uint(historyID))
	if err != nil {
		return errors.New("data tidak ditemukan")
	}

	bookNew, err := bookmodel.GetOneByID(historyNew.BookID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("buku tidak ditemukan")
	}
	if _, err := usermodel.GetOneByID(historyNew.UserID); errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user tidak ditemukan")
	}
	if _, err := statusborrowhistorymodel.GetOneByID(historyNew.StatusID); errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("status tidak ditemukan")
	}

	if historyNew.ReturnedAt.IsZero() {
		historyNew.StatusID = 1
	} else {
		historyNew.StatusID = 2
	}

	if err := database.DB.Model(&entities.BookBorrowHistory{}).Where("id = ?", historyID).Updates(map[string]interface{}{
		"book_id":     historyNew.BookID,
		"user_id":     historyNew.UserID,
		"status_id":   historyNew.StatusID,
		"borrowed_at": historyNew.BorrowedAt,
		"returned_at": historyNew.ReturnedAt,
		"keterangan":  historyNew.Keterangan,
	}).Error; err != nil {
		return err
	}

	if history.BookID != historyNew.BookID {
		book, err := bookmodel.GetOneByID(history.BookID)
		if err == nil {
			book.IsAvailable = true
			bookmodel.Update(book)
		}
	}

	if historyNew.StatusID == 1 {
		bookNew.IsAvailable = false
	} else {
		bookNew.IsAvailable = true
	}

	return bookmodel.Update(bookNew)
}

func Delete(historyId uint) error {
	history, err := GetOneByID(historyId)
	if err != nil {
		return err
	}

	book, _ := bookmodel.GetOneByID(history.BookID)

	if err = database.DB.Delete(&history).Error; err != nil {
		return err
	}

	book.IsAvailable = true
	return bookmodel.Update(book)
}