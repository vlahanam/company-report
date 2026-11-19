package business

import (
	"context"

	"company-report/internal/module/user/dto"
	"company-report/internal/module/user/entity"
)

type CreateUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, data map[string]interface{}, user *entity.User) (*entity.User, error)
}

type createUserBiz struct {
	repo CreateUserRepository
}

func NewCreateUserBiz(repo CreateUserRepository) *createUserBiz {
	return &createUserBiz{repo: repo}
}

func (biz *createUserBiz) CreateUser(ctx context.Context, req *dto.CreateUserReqDTO) (*entity.User, error) {
	var user *entity.User
	user, _ = biz.repo.GetUser(ctx, map[string]interface{}{"email": *req.Email}, user)
	if user != nil {
		return nil, nil
	}

	user = &entity.User{
		Email:        *req.Email,
		PasswordHash: *req.Password,
		FullName:     req.FullName,
		IsActive:     true,
		TokenVersion: 1,
	}

	if err := biz.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
