package usecase

import (
	domain "NurtureNest/pkg/domain"
	interfaces "NurtureNest/pkg/repository/interface"
	services "NurtureNest/pkg/usecase/interface"
	"context"
)

type goalTrackingUseCase struct {
	goalTrackingRepo interfaces.GoalTrackingRepository
}

func NewGoalTrackingUseCase(repo interfaces.GoalTrackingRepository) services.GoalTrackingUseCase {
	return &goalTrackingUseCase{
		goalTrackingRepo: repo,
	}
}

func (c *goalTrackingUseCase) SetGoal(ctx context.Context, goal domain.SetGoal) (domain.SetGoal, error) {
	goal, err := c.goalTrackingRepo.SetGoal(ctx, goal)

	return goal, err
}
