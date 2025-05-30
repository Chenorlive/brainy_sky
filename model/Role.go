package model

import "github.com/gofrs/uuid"

type Role struct {
	Base
	Name        string  `json:"name" `
	Description *string `json:"description"`
}

type RolePermission struct {
	Base
	RoleID       uuid.UUID   `json:"role_id" `
	PermissionID uuid.UUID   `json:"permission_id" `
	Role         *Role       `json:"role" gorm:"foreignKey:RoleID"`
	Permission   *Permission `json:"permission" gorm:"foreignKey:PermissionID"`
}

type Permission struct {
	Base
	Name        string  `json:"name" `
	Description *string `json:"description"`
}

type UserRole struct {
	Base
	UserID uuid.UUID `json:"user_id" `
	RoleID uuid.UUID `json:"role_id" `
	User   *User     `json:"user" gorm:"foreignKey:UserID"`
	Role   *Role     `json:"role" gorm:"foreignKey:RoleID"`
}
