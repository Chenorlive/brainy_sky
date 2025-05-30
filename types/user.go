package types

import "github.com/Chenorlive/brainy/model"

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterUser struct {
	FirstName   string       `json:"first_name" binding:"required"`
	LastName    string       `json:"last_name" binding:"required"`
	MiddleName  *string      `json:"middle_name"`
	Username    string       `json:"username" binding:"required"`
	Email       *string      `json:"email"`
	Password    string       `json:"password" binding:"required"`
	Phone       *string      `json:"phone"`
	LoginHint   *string      `json:"login_hint"`
	DataOfBirth *string      `json:"date_of_birth"`
	Address     *string      `json:"address"`
	Image       *string      `json:"image"`
	IsActive    bool         `json:"is_active" binding:"required"`
	UserRoles   []model.Role `json:"user_roles" binding:"required"`
}
