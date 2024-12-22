package usermodel

import (
	"github.com/rizqullorayhan/go-fiber-gorm/database"
	"github.com/rizqullorayhan/go-fiber-gorm/dto"
	"github.com/rizqullorayhan/go-fiber-gorm/entities"
	"github.com/rizqullorayhan/go-fiber-gorm/model/rolemodel"
	"golang.org/x/crypto/bcrypt"
	// "gorm.io/gorm"
)

func GetAll() ([]dto.GetUser, error) {
	var users []entities.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	var userDTOs []dto.GetUser
    for _, user := range users {
        userDTO := dto.GetUser{
            ID:      user.ID,
            Name:    user.Name,
            Email:   user.Email,
            Address: user.Address,
            Phone:   user.Phone,
            Role:   user.Role.Name,
        }
        userDTOs = append(userDTOs, userDTO)
    }
	return userDTOs, nil
}

func GetOneByID(userId uint) (*entities.User, error) {
	var user entities.User

	if err := database.DB.Preload("Role").First(&user, "id = ?", userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	if err := database.DB.Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func Create(user *dto.Register) error {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	role, _ := rolemodel.GetOneByName("User")
	newUser := entities.User{
		Name: user.Name,
		Email: user.Email,
		Password: string(hashPassword),
		Address: user.Address,
		Phone: user.Phone,
		RoleID: role.ID,
	}

	return database.DB.Create(&newUser).Error
}

func Update(user *entities.User) error {
	return database.DB.Save(user).Error
}

func UpdateEmail(email *dto.UpdateUserEmail, userId uint) error {
	return database.DB.Model(&entities.User{}).Where("id = ?", userId).Update("email", email.Email).Error
}

func UpdateRole(userId uint, roleId uint) error {
	return database.DB.Model(&entities.User{}).Where("id = ?", userId).Update("role_id", roleId).Error
}

func Delete(userId uint) error {
	return database.DB.Where("id = ?", userId).Delete(&entities.User{}).Error
}