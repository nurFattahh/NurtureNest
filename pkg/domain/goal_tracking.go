package domain

import "time"

type GoalTracking struct {
	Goal   time.Time `json:"goal"`
	UserID uint      `json:"user_id"`
}
