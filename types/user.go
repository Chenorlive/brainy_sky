package types

import (
	"github.com/gofrs/uuid"
)

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterUser struct {
	FirstName   string  `json:"first_name" binding:"required"`
	LastName    string  `json:"last_name" binding:"required"`
	MiddleName  *string `json:"middle_name"`
	Username    string  `json:"username" binding:"required"`
	Email       *string `json:"email"`
	Password    string  `json:"password" binding:"required"`
	Phone       *string `json:"phone"`
	DateOfBirth *string `json:"date_of_birth"`
	Address     *string `json:"address"`
	Image       *string `json:"image"`
	IsActive    bool    `json:"is_active" binding:"required"`
	Role        string  `json:"user_roles" binding:"required"`
}

type UpdateUser struct {
	ID          uuid.UUID `json:"id" binding:"required"`
	FirstName   *string   `json:"first_name" binding:"required"`
	LastName    *string   `json:"last_name" binding:"required"`
	MiddleName  *string   `json:"middle_name"`
	Username    *string   `json:"username" binding:"required"`
	Email       *string   `json:"email"`
	Phone       *string   `json:"phone"`
	DateOfBirth *string   `json:"date_of_birth"`
	Address     *string   `json:"address"`
	Image       *string   `json:"image"`
}

type UpdatePassword struct {
	ID          uuid.UUID `json:"id" binding:"required"`
	OldPassword string    `json:"old_password" binding:"required"`
	NewPassword string    `json:"new_password" binding:"required"`
}
