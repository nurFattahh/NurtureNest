package repository

import (
	domain "NurtureNest/pkg/domain"
	interfaces "NurtureNest/pkg/repository/interface"
	"context"

	"gorm.io/gorm"
)

type goalTrackingDatabase struct {
	DB *gorm.DB
}

func NewGoalTrackingRepository(DB *gorm.DB) interfaces.GoalTrackingRepository {
	return &goalTrackingDatabase{DB}
}

func (c *goalTrackingDatabase) SetGoal(ctx context.Context, goal domain.SetGoal) (domain.SetGoal, error) {
	var goalTrack domain.GoalTracking
	err := c.DB.Model(goalTrack).Create(&goal).Error

	return goal, err
}
