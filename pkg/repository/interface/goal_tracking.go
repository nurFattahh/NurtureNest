package interfaces

import (
	"NurtureNest/pkg/domain"
	"context"
)

type GoalTrackingRepository interface {
	SetGoal(ctx context.Context, goal domain.SetGoal) (domain.SetGoal, error)
}
