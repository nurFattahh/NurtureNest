package interfaces

import (
	"context"

	"NurtureNest/pkg/domain"
)

type TutorRepository interface {
	Save(ctx context.Context, user domain.Tutor) (domain.Tutor, error)
}
