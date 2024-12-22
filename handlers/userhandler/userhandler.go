package userhandler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rizqullorayhan/go-fiber-gorm/dto"
	"github.com/rizqullorayhan/go-fiber-gorm/model/usermodel"
	"github.com/rizqullorayhan/go-fiber-gorm/validators"
)

func GetAll(ctx *fiber.Ctx) error {
	users, err := usermodel.GetAll()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.JSON(users)
}

func Create(ctx *fiber.Ctx) error {
	user := new(dto.Register)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if errorValidator := validators.ValidateStruct(user); errorValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidator,
		})
	}

	if err := usermodel.Create(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil menambahkan data.",
	})
}

func GetById(ctx *fiber.Ctx) error {
	userIdParam := ctx.Params("id")
	userId, err := strconv.ParseUint(userIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	user, err := usermodel.GetOneByID(uint(userId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan.",
		})
	}

	row := dto.GetUser{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
		Role:    user.Role.Name,
	}

	return ctx.Status(fiber.StatusOK).JSON(row)
}

func Update(ctx *fiber.Ctx) error {
	userIdParam := ctx.Params("id")
	userId, err := strconv.ParseUint(userIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	// Mendapatkan data dari body request
	userNew := new(dto.UpdateUser)
	if err := ctx.BodyParser(userNew); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Validasi data
	if errorValidator := validators.ValidateStruct(userNew); errorValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidator,
		})
	}

	// Cek apakah data ada atau tidak ada
	user, err := usermodel.GetOneByID(uint(userId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan.",
		})
	}

	// Ubah nilai field yang ingin diupdate
	user.Name = userNew.Name
	user.Address = userNew.Address
	user.Phone = userNew.Phone

	// Simpan perubahan
	if err := usermodel.Update(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengupdate data.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil mengupdate data.",
	})
}

func UpdateEmail(ctx *fiber.Ctx) error {
	userIdParam := ctx.Params("id")
	userId, err := strconv.ParseUint(userIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	email := new(dto.UpdateUserEmail)
	if err := ctx.BodyParser(email); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if errorValidator := validators.ValidateStruct(email); errorValidator != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidator,
		})
	}

	if err := usermodel.UpdateEmail(email, uint(userId)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengupdate data.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil mengupdate email.",
	})
}

func UpdateRole(ctx *fiber.Ctx) error {
	userIdParam := ctx.Params("id")
	userId, err := strconv.ParseUint(userIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	role := new(dto.UpdateUserRole)
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

	if err := usermodel.UpdateRole(role.RoleID, uint(userId)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengupdate data.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil mengupdate email.",
	})
}

func Delete(ctx *fiber.Ctx) error {
	userIdParam := ctx.Params("id")
	userId, err := strconv.ParseUint(userIdParam, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID parameter.",
		})
	}

	if err := usermodel.Delete(uint(userId)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Berhasil menghapus data.",
	})
}
