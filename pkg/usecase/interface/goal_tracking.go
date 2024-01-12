package interfaces

import (
	domain "NurtureNest/pkg/domain"
	"context"
)

type GoalTrackingUseCase interface {
	SetGoal(ctx context.Context, goal domain.SetGoal) (domain.SetGoal, error)
}
