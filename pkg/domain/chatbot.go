package domain

import "time"

type ChatBot struct {
	CreatedAt time.Time `json:"created_at"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	UserID    uint      `json:"-"`
}
