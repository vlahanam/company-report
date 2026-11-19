package repository

import (
	"context"

	"company-report/internal/module/user/entity"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

func (repo *userRepo) CreateUser(ctx context.Context, user *entity.User) error {
	if err := repo.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	return nil
}
