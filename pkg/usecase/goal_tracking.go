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

func (c *goalTrackingUseCase) SetGoal(ctx context.Context, goal domain.GoalTracking, user_id uint) (domain.GoalTracking, error) {
	goal, err := c.goalTrackingRepo.SetGoal(ctx, goal, user_id)

	return goal, err
}

func (c *goalTrackingUseCase) SaveProgress(ctx context.Context, progress domain.RequestGoalResult, user_id uint) error {
	err := c.goalTrackingRepo.SaveProgress(ctx, progress, user_id)
	return err
}

func (c *goalTrackingUseCase) GetGoalTracking(ctx context.Context, user_id uint) (domain.GoalTracking, error) {
	goalTrack, err := c.goalTrackingRepo.GetGoalTracking(ctx, user_id)
	return goalTrack, err
}

func (c *goalTrackingUseCase) GoalResult(ctx context.Context, goalTrack domain.GoalTracking, user_id uint, status string, result bool) (domain.GoalResult, error) {
	goalResult, err := c.goalTrackingRepo.GoalResult(ctx, goalTrack, user_id, status, result)

	return goalResult, err
}
