package rolehandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rizqullorayhan/go-fiber-gorm/dto"
	"github.com/rizqullorayhan/go-fiber-gorm/model/rolemodel"
	"github.com/rizqullorayhan/go-fiber-gorm/validators"
)

func GetAll(ctx *fiber.Ctx) error {
	categories, err := rolemodel.GetAll()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.JSON(categories)
}

func Create(ctx *fiber.Ctx) error {
	role := new(dto.CreateRole)
	if err := ctx.BodyParser(role); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if errorValidator := validators.ValidateStruct(role); errorValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidator,
		})
	}

	if err := rolemodel.Create(role); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menambahkan role.",
	})
}

func GetById(ctx *fiber.Ctx) error {
	roleIdParam := ctx.Params("id")
	roleId, err := strconv.ParseUint(roleIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	role, err := rolemodel.GetOneByID(uint(roleId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan.",
		})
	}

	row := dto.GetRole{
		ID:    role.ID,
		Name:  role.Name,
	}

	return ctx.Status(fiber.StatusOK).JSON(row)
}

func Update(ctx *fiber.Ctx) error {
	roleIdParam := ctx.Params("id")
	roleId, err := strconv.ParseUint(roleIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	// Mendapatkan data dari body request
	roleNew := new(dto.UpdateRole)
	if err := ctx.BodyParser(roleNew); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Validasi data
	if errorValidator := validators.ValidateStruct(roleNew); errorValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidator,
		})
	}

	// Cek apakah data ada atau tidak ada
	role, err := rolemodel.GetOneByID(uint(roleId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan.",
		})
	}

	// Ubah nilai field yang ingin diupdate
	role.Name = roleNew.Name

	// Simpan perubahan
	if err := rolemodel.Update(role); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengupdate data.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil mengupdate data.",
	})
}

func Delete(ctx *fiber.Ctx) error {
	roleIdParam := ctx.Params("id")
	roleId, err := strconv.ParseUint(roleIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	if err := rolemodel.Delete(uint(roleId)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil menghapus data.",
	})
}