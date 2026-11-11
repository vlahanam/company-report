package business

import (
	"context"
	"errors"

	"company-report/internal/module/user/dto"
	"company-report/internal/module/user/entity"
)

type CreateUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
}

type createUserBiz struct {
	repo CreateUserRepository
}

func NewCreateUserBiz(repo CreateUserRepository) *createUserBiz {
	return &createUserBiz{repo: repo}
}

func (biz *createUserBiz) CreateUser(ctx context.Context, req *dto.CreateUserReqDTO) (*entity.User, error) {
	if req == nil || req.Email == nil || req.Password == nil {
		return nil, errors.New("email and password are required")
	}

	user := entity.User{
		Email:        *req.Email,
		PasswordHash: *req.Password,
		FullName:     req.FullName,
		IsActive:     true,
		TokenVersion: 1,
	}

	if err := biz.repo.CreateUser(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
