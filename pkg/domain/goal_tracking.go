package domain

import "time"

type GoalTracking struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Goal      string    `json:"goal"`
	UserID    uint      `json:"user_id"`
	LimitTime time.Time `json:"limit"`
	Status    string    `json:"status"`
	IsActive  bool      `json:"isActive"`
	Progress  string    `json:"progress"`
	Result    bool      `json:"-"`
}

type RequestSetGoal struct {
	GoalHour   int `json:"goal_hour"`
	GoalMinute int `json:"goal_minute"`
	GoalSecond int `json:"goal_second"`
}

type GoalResult struct {
	ID        uint      `json:"id"`
	Progress  string    `json:"progress"`
	Goal      string    `json:"goal"`
	UserID    uint      `json:"user_id"`
	LimitTime time.Time `json:"limit"`
	Status    string    `json:"status"`
	Result    bool      `json:"result"`
}

type RequestGoalResult struct {
	Progress string `json:"progress"`
}
