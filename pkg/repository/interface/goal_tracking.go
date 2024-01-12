package interfaces

import (
	"NurtureNest/pkg/domain"
	"context"
)

type TutorRepository12 interface {
	Save(ctx context.Context, user domain.Tutor) (domain.Tutor, error)
}
