package repository

import (
	"company-report/internal/module/user/entity"
	"context"
)

func (repo *userRepo) GetUser(ctx context.Context, data map[string]interface{}, user *entity.User) (*entity.User, error) {
	if err := repo.db.WithContext(ctx).Where(data).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
