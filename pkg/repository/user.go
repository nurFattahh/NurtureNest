package repository

import (
	"context"

	"NurtureNest/pkg/api/sdk/crypto"
	domain "NurtureNest/pkg/domain"
	interfaces "NurtureNest/pkg/repository/interface"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

// fungsi register untuk user
func (c *userDatabase) UserRegister(ctx context.Context, model domain.UserRegister) (*domain.Users, error) {

	hashPassword, err := crypto.HashValue(model.Password)
	if err != nil {
		return nil, err
	}

	var user domain.Users = domain.Users{
		FullName: model.FullName,
		Email:    model.Email,
		Username: model.Username,
		Password: hashPassword,
		Gender:   model.Gender,
		Role:     model.Role,
	}

	result := c.DB.Create(&user)
	if result.Error != nil {
		return nil, err
	}

	return &user, nil
}

func (c *userDatabase) UserFindByUsernameOrEmail(ctx context.Context, username string, email string) (domain.Users, error) {
	var user domain.Users
	err := c.DB.Where("username = ? or email = ?", username, email).First(&user).Error

	return user, err
}

func (c *userDatabase) UserFindAll(ctx context.Context) ([]domain.Users, error) {
	var users []domain.Users
	err := c.DB.Find(&users).Error
	return users, err
}

func (c *userDatabase) FindAll(ctx context.Context) ([]domain.Users, error) {
	var users []domain.Users
	err := c.DB.Find(&users).Error

	return users, err
}

func (c *userDatabase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	var user domain.Users
	err := c.DB.First(&user, id).Error

	return user, err
}

func (c *userDatabase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := c.DB.Save(&user).Error

	return user, err
}

func (c *userDatabase) Delete(ctx context.Context, user domain.Users) error {
	err := c.DB.Delete(&user).Error

	return err
}
