package repository

import (
	domain "NurtureNest/pkg/domain"
	interfaces "NurtureNest/pkg/repository/interface"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type goalTrackingDatabase struct {
	DB *gorm.DB
}

func NewGoalTrackingRepository(DB *gorm.DB) interfaces.GoalTrackingRepository {
	return &goalTrackingDatabase{DB}
}

func (c *goalTrackingDatabase) SetGoal(ctx context.Context, goal domain.GoalTracking, user_id uint) (domain.GoalTracking, error) {
	var goalTrack domain.GoalTracking
	_ = c.DB.Model(goalTrack).Where("user_id =?", user_id).Where("is_active =?", true).Update("is_active", false).Error
	err := c.DB.Model(goalTrack).Create(&goal).Error

	return goal, err
}

func (c *goalTrackingDatabase) SaveProgress(ctx context.Context, progress domain.RequestGoalResult, user_id uint) error {
	err := c.DB.Model(domain.GoalTracking{}).Where("user_id =? and is_active =?", user_id, true).Update("progress", progress.Progress).Error
	return err
}

func (c *goalTrackingDatabase) GetGoalTracking(ctx context.Context, user_id uint) (domain.GoalTracking, error) {
	var user domain.GoalTracking
	err := c.DB.Model(user).Where("is_active =? and user_id =?", true, user_id).Find(&user).Error
	fmt.Println(user.LimitTime)
	return user, err
}

func (c *goalTrackingDatabase) GoalResult(ctx context.Context, goalTrack domain.GoalTracking, user_id uint, status string, result bool) (domain.GoalResult, error) {
	var finalResult domain.GoalResult = domain.GoalResult{
		ID:        goalTrack.ID,
		Progress:  goalTrack.Progress,
		Goal:      goalTrack.Goal,
		UserID:    goalTrack.UserID,
		LimitTime: goalTrack.LimitTime,
		Status:    status,
	}

	var model domain.GoalTracking

	_ = c.DB.Model(model).Where("user_id =? and is_active =?", user_id, true).Updates(domain.GoalTracking{Progress: goalTrack.Progress, Status: status, Result: result})
	err := c.DB.Model(goalTrack).Where("user_id =? and is_active =?", user_id, true).Last(&finalResult).Error

	return finalResult, err
}
