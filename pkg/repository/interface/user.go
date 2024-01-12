package interfaces

import (
	"context"

	"NurtureNest/pkg/domain"
)

type UserRepository interface {
	UserRegister(ctx context.Context, model domain.UserRegister) (*domain.Users, error)
	UserFindByUsernameOrEmail(ctx context.Context, username string, email string) (domain.Users, error)
	UserFindAll(ctx context.Context) ([]domain.Users, error)
	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Delete(ctx context.Context, user domain.Users) error
}
