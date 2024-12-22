package bookhandler

import (
	// "fmt"
	// "image"
	// _ "image/jpeg"
	// _ "image/png"

	"mime/multipart"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rizqullorayhan/go-fiber-gorm/dto"
	"github.com/rizqullorayhan/go-fiber-gorm/model/bookmodel"
	"github.com/rizqullorayhan/go-fiber-gorm/utils"
	"github.com/rizqullorayhan/go-fiber-gorm/validators"
)

func GetAll(ctx *fiber.Ctx) error {
	books, err := bookmodel.GetAll()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.JSON(books)
}

func Create(ctx *fiber.Ctx) error {
	book := new(dto.CreateBook)
	if err := ctx.BodyParser(book); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
        })
	}

	// Parsing categories sebagai array
	categoryIDs := ctx.FormValue("categories") // Mendapatkan kategori sebagai string
    if categoryIDs == "" {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Categories are required.",
        })
    }

	// Konversi categoryIDs menjadi array integer
    ids := strings.Split(categoryIDs, ",") // Jika dikirim dalam format CSV
    for _, id := range ids {
        categoryID, err := strconv.Atoi(strings.TrimSpace(id))
        if err != nil {
            return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "message": "Invalid category ID format.",
            })
        }
        book.Categories = append(book.Categories, categoryID)
    }

	if errorValidator := validators.ValidateStruct(book); errorValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidator,
		})
	}

	file, err := ctx.FormFile("cover")
	if err != nil || file == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "File cover tidak ditemukan.",
		})
	}

	book.Cover, err = uploadFile(file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := bookmodel.Create(book); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menambahkan buku.",
	})
}

func GetById(ctx *fiber.Ctx) error {
	bookIdParam := ctx.Params("id")
	bookId, err := strconv.ParseUint(bookIdParam, 10, 32)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid ID parameter.",
        })
    }

	book, err := bookmodel.GetOneByID(uint(bookId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan.",
		})
	}

	row := dto.GetBook{
		ID: book.ID,
		Title: book.Title,
		Author: book.Author,
		Cover: book.Cover,
		IsAvailable: book.IsAvailable,
		Categories: book.Categories,
	}

	return ctx.Status(fiber.StatusOK).JSON(row)
}

func Update(ctx *fiber.Ctx) error {
	bookIdParam := ctx.Params("id")
	bookId, err := strconv.ParseUint(bookIdParam, 10, 32)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid ID parameter.",
        })
    }

	// Mendapatkan data dari body request
	bookNew := new(dto.UpdateBook)
	if err := ctx.BodyParser(bookNew); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
        })
	}

	// Validasi data
	if errorValidator := validators.ValidateStruct(bookNew); errorValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidator,
		})
	}

	// Cek apakah data ada atau tidak ada
	book, err := bookmodel.GetOneByID(uint(bookId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan.",
		})
	}

	file, err := ctx.FormFile("cover")
	if err == nil || file != nil {
		newCover, err := uploadFile(file)
		if err == nil {
			_ = utils.DeleteFile(book.Cover)
			book.Cover = newCover
		}
	}

	// Ubah nilai field yang ingin diupdate
	book.Title = bookNew.Title
	book.Author = bookNew.Author
	book.IsAvailable = bookNew.IsAvailable

	// Simpan perubahan
	if err := bookmodel.UpdateWithCategory(book, bookNew.Categories); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengupdate data.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil mengupdate data.",
	})
}

func Delete(ctx *fiber.Ctx) error {
	bookIdParam := ctx.Params("id")
	bookId, err := strconv.ParseUint(bookIdParam, 10, 32)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid ID parameter.",
        })
    }

	// Cek apakah data ada atau tidak ada
	book, err := bookmodel.GetOneByID(uint(bookId))
	if err != nil {
		return ctx.JSON(fiber.Map{
			"message": "Berhasil menghapus data.",
		})
	}

	if err := bookmodel.Delete(uint(bookId)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data.",
		})
	}

	_ = utils.DeleteFile(book.Cover)

	return ctx.JSON(fiber.Map{
		"message": "Berhasil menghapus data.",
	})
}

func uploadFile(file *multipart.FileHeader) (string, error) {
	err := validators.ValidateImage(file)
	if err != nil {
		return "", err
	}

	destFolder := "/uploads/cover/"
	if err := utils.SaveFile(file, destFolder, file.Filename); err != nil {
		return "", err
	}
	return destFolder+file.Filename, nil
}