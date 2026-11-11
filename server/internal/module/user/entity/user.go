package entity

import "company-report/internal/common"

type User struct {
	common.BaseEntity
	Email        string  `json:"email" gorm:"column:email"`
	PasswordHash string  `json:"-" gorm:"column:password_hash"`
	FullName     *string `json:"full_name,omitempty" gorm:"column:full_name"`
	IsActive     bool    `json:"is_active" gorm:"column:is_active"`
	TokenVersion int     `json:"token_version" gorm:"column:token_version"`
}

func (User) TableName() string {
	return "users"
}
