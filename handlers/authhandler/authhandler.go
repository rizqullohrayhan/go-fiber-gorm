package authhandler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rizqullorayhan/go-fiber-gorm/dto"
	"github.com/rizqullorayhan/go-fiber-gorm/model/usermodel"
	"github.com/rizqullorayhan/go-fiber-gorm/utils"
	"github.com/rizqullorayhan/go-fiber-gorm/validators"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *fiber.Ctx) error {
	credentials := new(dto.Login)
	if err := ctx.BodyParser(credentials); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
        })
	}

	if err := validators.ValidateStruct(credentials); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	user, err := usermodel.FindByEmail(credentials.Email)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Email tidak terdaftar.",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err !=nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Password salah.",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  user.Name,
		"role":  user.Role.Name,
		"exp":   time.Now().Add(time.Hour).Unix(),
	}

	t, err := utils.GenerateToken(&claims)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{
		"message": "Login Berhasil",
		"token": t,
	})
}

func Register(ctx *fiber.Ctx) error {
	user := new(dto.Register)
	if err := ctx.BodyParser(user); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid request body",
        })
	}

	if err := validators.ValidateStruct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := usermodel.Create(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Registrasi Berhasil.",
	})
}