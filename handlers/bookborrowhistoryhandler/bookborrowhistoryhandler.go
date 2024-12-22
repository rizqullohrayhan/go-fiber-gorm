package bookborrowhistoryhandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rizqullorayhan/go-fiber-gorm/dto"
	"github.com/rizqullorayhan/go-fiber-gorm/model/bookborrowhistorymodel"
	"github.com/rizqullorayhan/go-fiber-gorm/validators"
)

func GetAll(ctx *fiber.Ctx) error {
	history, err := bookborrowhistorymodel.GetAll()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.JSON(history)
}

func GetByUser(ctx *fiber.Ctx) error {
	userIdParam := ctx.Params("userId")
	userId, err := strconv.ParseUint(userIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}
	history, err := bookborrowhistorymodel.GetByUser(uint(userId))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.JSON(history)
}

func GetByBook(ctx *fiber.Ctx) error {
	bookIdParam := ctx.Params("bookId")
	bookId, err := strconv.ParseUint(bookIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}
	history, err := bookborrowhistorymodel.GetByBook(uint(bookId))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.JSON(history)
}

func GetByUserAndBook(ctx *fiber.Ctx) error {
	bookIdParam := ctx.Params("bookId")
	bookId, err := strconv.ParseUint(bookIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}
	userIdParam := ctx.Params("userId")
	userId, err := strconv.ParseUint(userIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}
	history, err := bookborrowhistorymodel.GetByUserAndBook(uint(userId), uint(bookId))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.JSON(history)
}

func GetById(ctx *fiber.Ctx) error {
	historyIdParam := ctx.Params("id")
	historyId, err := strconv.ParseUint(historyIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}
	history, err := bookborrowhistorymodel.GetOneByID(uint(historyId))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	if history == nil {
		// Data tidak ditemukan
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "History not found.",
		})
	}

	return ctx.JSON(history)
}

func GetLastBorrowedBook(ctx *fiber.Ctx) error {
	userIdParam := ctx.Params("userId")
	userId, err := strconv.ParseUint(userIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	books, err := bookborrowhistorymodel.GetUserLatestBorrowBook(uint(userId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.JSON(books)
}

func Create(ctx *fiber.Ctx) error {
	bookBorrowHistory := new(dto.CreateBookBorrowHistory)
	if err := ctx.BodyParser(bookBorrowHistory); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if errorValidator := validators.ValidateStruct(bookBorrowHistory); errorValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidator,
		})
	}

	if err := bookborrowhistorymodel.Create(bookBorrowHistory); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menambahkan riwayat peminjaman buku.",
	})
}

func Update(ctx *fiber.Ctx) error {
	historyIdParam := ctx.Params("id")
	historyId, err := strconv.ParseUint(historyIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	history := new(dto.UpdateBookBorrowHistory)
	if err := ctx.BodyParser(history); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			// "message": err,
		})
	}

	if errorValidator := validators.ValidateStruct(history); errorValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidator,
		})
	}

	// history.BookID = historyNew.BookID
	// history.UserID = historyNew.UserID
	// history.StatusID = historyNew.StatusID
	// history.BorrowedAt = historyNew.BorrowedAt
	// history.ReturnedAt = historyNew.ReturnedAt
	// history.Keterangan = historyNew.Keterangan

	if err := bookborrowhistorymodel.Update(uint(historyId), history); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengupdate history.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil mengupdate history.",
	})
}

func Delete(ctx *fiber.Ctx) error {
	historyIdParam := ctx.Params("id")
	historyId, err := strconv.ParseUint(historyIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	if err := bookborrowhistorymodel.Delete(uint(historyId)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil menghapus data.",
	})
}
