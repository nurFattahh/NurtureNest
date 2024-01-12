package usecase

import (
	"context"

	domain "NurtureNest/pkg/domain"
	interfaces "NurtureNest/pkg/repository/interface"
	services "NurtureNest/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) UserRegister(ctx context.Context, model domain.UserRegister) (*domain.Users, error) {
	user, err := c.userRepo.UserRegister(ctx, model)
	return user, err
}

func (c *userUseCase) UserFindByUsernameOrEmail(ctx context.Context, username string, email string) (domain.Users, error) {
	user, err := c.userRepo.UserFindByUsernameOrEmail(ctx, username, email)
	return user, err
}

func (c *userUseCase) UserFindAll(ctx context.Context) ([]domain.Users, error) {
	users, err := c.userRepo.UserFindAll(ctx)
	return users, err
}

func (c *userUseCase) FindAll(ctx context.Context) ([]domain.Users, error) {
	users, err := c.userRepo.FindAll(ctx)
	return users, err
}

func (c *userUseCase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	user, err := c.userRepo.FindByID(ctx, id)
	return user, err
}

func (c *userUseCase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	user, err := c.userRepo.Save(ctx, user)

	return user, err
}

func (c *userUseCase) Delete(ctx context.Context, user domain.Users) error {
	err := c.userRepo.Delete(ctx, user)

	return err
}
