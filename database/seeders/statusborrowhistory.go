package seeders

import (
	"fmt"

	"github.com/rizqullorayhan/go-fiber-gorm/database"
	"github.com/rizqullorayhan/go-fiber-gorm/entities"
)

func SeedStatusBorrowHistories() error {
	statues := []entities.StatusBorrowHistory{
		{Name: "Dipinjam"},
		{Name: "Dikembalikan"},
	}

	for _, status := range statues {
		// Gunakan Create atau FirstOrCreate untuk menghindari duplikasi
		if err := database.DB.FirstOrCreate(&entities.StatusBorrowHistory{}, status).Error; err != nil {
			return fmt.Errorf("failed to seed status borrow history %s: %w", status.Name, err)
		}
	}

	fmt.Println("Status borrow history seeding completed successfully.")
	return nil
}