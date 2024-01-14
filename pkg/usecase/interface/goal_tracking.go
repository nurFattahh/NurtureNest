package interfaces

import (
	domain "NurtureNest/pkg/domain"
	"context"
)

type GoalTrackingUseCase interface {
	SetGoal(ctx context.Context, goal domain.GoalTracking, user_id uint) (domain.GoalTracking, error)
	SaveProgress(ctx context.Context, progress domain.RequestGoalResult, user_id uint) error
	GetGoalTracking(ctx context.Context, user_id uint) (domain.GoalTracking, error)
	GoalResult(ctx context.Context, goalTrack domain.GoalTracking, user_id uint, status string, result bool) (domain.GoalResult, error)
}
