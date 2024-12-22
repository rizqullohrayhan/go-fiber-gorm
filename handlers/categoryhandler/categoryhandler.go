package categoryhandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rizqullorayhan/go-fiber-gorm/dto"
	"github.com/rizqullorayhan/go-fiber-gorm/model/categorymodel"
	"github.com/rizqullorayhan/go-fiber-gorm/validators"
)

func GetAll(ctx *fiber.Ctx) error {
	categories, err := categorymodel.GetAll()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.JSON(categories)
}

func Create(ctx *fiber.Ctx) error {
	category := new(dto.CreateCategory)
	if err := ctx.BodyParser(category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if errorValidator := validators.ValidateStruct(category); errorValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidator,
		})
	}

	if err := categorymodel.Create(category); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menambahkan kategori.",
	})
}

func GetById(ctx *fiber.Ctx) error {
	categoryIdParam := ctx.Params("id")
	categoryId, err := strconv.ParseUint(categoryIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	book, err := categorymodel.GetOneByID(uint(categoryId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan.",
		})
	}

	row := dto.GetCategory{
		ID:    book.ID,
		Name:  book.Name,
		Books: book.Books,
	}

	return ctx.Status(fiber.StatusOK).JSON(row)
}

func Update(ctx *fiber.Ctx) error {
	categoryIdParam := ctx.Params("id")
	categoryId, err := strconv.ParseUint(categoryIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	// Mendapatkan data dari body request
	categoryNew := new(dto.UpdateCategory)
	if err := ctx.BodyParser(categoryNew); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Validasi data
	if errorValidator := validators.ValidateStruct(categoryNew); errorValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidator,
		})
	}

	// Cek apakah data ada atau tidak ada
	category, err := categorymodel.GetOneByID(uint(categoryId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan.",
		})
	}

	// Ubah nilai field yang ingin diupdate
	category.Name = categoryNew.Name

	// Simpan perubahan
	if err := categorymodel.Update(category); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengupdate data.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil mengupdate data.",
	})
}

func Delete(ctx *fiber.Ctx) error {
	categoryIdParam := ctx.Params("id")
	categoryId, err := strconv.ParseUint(categoryIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	if err := categorymodel.Delete(uint(categoryId)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil menghapus data.",
	})
}
