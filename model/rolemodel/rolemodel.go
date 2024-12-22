package rolemodel

import (
	"github.com/rizqullorayhan/go-fiber-gorm/database"
	"github.com/rizqullorayhan/go-fiber-gorm/dto"
	"github.com/rizqullorayhan/go-fiber-gorm/entities"
	"gorm.io/gorm"
)

func GetAll() ([]dto.GetRole, error) {
	var roles []entities.Role
	if err := database.DB.Find(&roles).Error; err != nil {
		return nil, err
	}

	var roleDTOs []dto.GetRole
	for _, role := range roles {
		roleDTO := dto.GetRole{
			ID:   role.ID,
			Name: role.Name,
		}
		roleDTOs = append(roleDTOs, roleDTO)
	}
	return roleDTOs, nil
}

func GetOneByID(roleId uint) (*entities.Role, error) {
	var role entities.Role

	if err := database.DB.Where("id = ?", roleId).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func GetOneByName(roleName string) (*entities.Role, error) {
	var role entities.Role

	if err := database.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func Create(role *dto.CreateRole) error {
	newRole := entities.Role{
		Name: role.Name,
	}

	// Cari data yang sudah dihapus (soft deleted) dengan email yang sama
    result := database.DB.Unscoped().Where("name = ?", newRole.Name).First(&newRole)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            // Jika tidak ada data, buat record baru
            return database.DB.Create(newRole).Error
        }
        return result.Error
    }

	newRole.DeletedAt = gorm.DeletedAt{}
	return database.DB.Save(&newRole).Error
}

func Update(role *entities.Role) error {
	return database.DB.Save(role).Error
}

func Delete(userId uint) error {
	return database.DB.Where("id = ?", userId).Delete(&entities.Role{}).Error
}