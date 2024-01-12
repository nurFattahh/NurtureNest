package interfaces

import (
	"context"

	domain "NurtureNest/pkg/domain"
)

type TutorUseCase interface {
	Save(ctx context.Context, user domain.Tutor) (domain.Tutor, error)
}
