package domain

type Tutor struct {
	ID      uint   `json:"id" gorm:"unique;not null"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Koin    int    `json:"koin"`
}
