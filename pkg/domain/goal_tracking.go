package domain

import "time"

type GoalTracking struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Goal      time.Time `json:"goal"`
	Progress  time.Time `json:"progress"`
	UserID    uint      `json:"user_id"`
	LimitTime time.Time `json:"limit"`
}

type SetGoal struct {
	Goal      time.Time `json:"goal"`
	UserID    uint      `json:"user_id"`
	LimitTime time.Time `json:"limit"`
}

type RequestSetGoal struct {
	GoalHour   int `json:"goal_hour"`
	GoalMinute int `json:"goal_minute"`
	GoalSecond int `json:"goal_second"`
}

type CheckProgress struct {
	Goal      time.Time `json:"goal"`
	Progress  time.Time `json:"progress"`
	UserID    uint      `json:"user_id"`
	LimitTime time.Time `json:"limit"`
}
