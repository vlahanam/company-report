package repository

import (
	"context"

	"gorm.io/gorm"
)

type getUserRepo struct {
	db *gorm.DB
}

func NewGetUserRepo(db *gorm.DB) *createUserRepo {
	return &createUserRepo{db: db}
}

func (repo *createUserRepo) GetUser(ctx context.Context, data map[string]interface{}) error {
	if err := repo.db.WithContext(ctx).Where(data).Error; err != nil {
		return err
	}
	return nil
}
