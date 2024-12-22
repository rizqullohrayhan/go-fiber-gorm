package seeders

import (
	"fmt"

	"github.com/rizqullorayhan/go-fiber-gorm/database"
	"github.com/rizqullorayhan/go-fiber-gorm/entities"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() error {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), 8)
	users := []entities.User{
		{
			Name: "Rizqulloh Rayhan",
			Email: "rizqullohferdiansyah41@gmail.com",
			Password: string(hashPassword),
			Address: "Jombang, Jawa Timur",
			Phone: "0123456789",
			RoleID: 1,
		},
		// {
		// 	Name: "Rayhan Ferdiansyah",
		// 	Email: "rizqullohrayhan@student.uns.ac.id",
		// 	Password: string(hashPassword),
		// 	Address: "Jombang, Jawa Timur",
		// 	Phone: "0123456789",
		// 	RoleID: 2,
		// },
	}

	for _, user := range users {
		// Gunakan Create atau FirstOrCreate untuk menghindari duplikasi
		if err := database.DB.Where(entities.User{Email: user.Email}).Attrs(entities.User{Email: user.Email}).FirstOrCreate(&entities.User{}, user).Error; err != nil {
			return fmt.Errorf("failed to seed user %s: %w", user.Name, err)
		}
	}

	fmt.Println("User seeding completed successfully.")
	return nil
}