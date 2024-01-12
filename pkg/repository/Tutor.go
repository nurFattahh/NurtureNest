package repository

import (
	"context"

	domain "NurtureNest/pkg/domain"
	interfaces "NurtureNest/pkg/repository/interface"

	"gorm.io/gorm"
)

type tutorDatabase struct {
	DB *gorm.DB
}

func NewTutorRepository(DB *gorm.DB) interfaces.TutorRepository {
	return &tutorDatabase{DB}
}

func (c *tutorDatabase) Save(ctx context.Context, user domain.Tutor) (domain.Tutor, error) {
	err := c.DB.Save(&user).Error

	return user, err
}
