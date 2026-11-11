package repository

import (
	"context"

	"company-report/internal/module/user/entity"

	"gorm.io/gorm"
)

type createUserRepo struct {
	db *gorm.DB
}

func NewCreateUserRepo(db *gorm.DB) *createUserRepo {
	return &createUserRepo{db: db}
}

func (repo *createUserRepo) CreateUser(ctx context.Context, user *entity.User) error {
	if err := repo.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}
