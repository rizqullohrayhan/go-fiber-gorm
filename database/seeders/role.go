package seeders

import (
	"fmt"

	"github.com/rizqullorayhan/go-fiber-gorm/database"
	"github.com/rizqullorayhan/go-fiber-gorm/entities"
)

func SeedRoles() error {
	roles := []entities.Role{
		{Name: "Admin"},
		{Name: "User"},
	}

	for _, role := range roles {
		// Gunakan Create atau FirstOrCreate untuk menghindari duplikasi
		if err := database.DB.FirstOrCreate(&entities.Role{}, role).Error; err != nil {
			return fmt.Errorf("failed to seed role %s: %w", role.Name, err)
		}
	}

	fmt.Println("Role seeding completed successfully.")
	return nil
}