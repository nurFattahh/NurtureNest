package usecase

import (
	"context"

	domain "NurtureNest/pkg/domain"
	interfaces "NurtureNest/pkg/repository/interface"
	services "NurtureNest/pkg/usecase/interface"
)

type tutorUseCase struct {
	tutorRepo interfaces.TutorRepository
}

func NewTutorUseCase(repo interfaces.TutorRepository) services.TutorUseCase {
	return &tutorUseCase{
		tutorRepo: repo,
	}
}

func (c *tutorUseCase) Save(ctx context.Context, user domain.Tutor) (domain.Tutor, error) {
	user, err := c.tutorRepo.Save(ctx, user)

	return user, err
}
