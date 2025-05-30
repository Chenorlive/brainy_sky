package types

import "github.com/gofrs/uuid"

type NewRole struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

type UpdateRole struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Name        *string   `json:"name" validate:"required"`
	Description *string   `json:"description"`
}

type NewPermission struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

type UpdatePermission struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Name        *string   `json:"name" validate:"required"`
	Description *string   `json:"description"`
}

type NewRolePermission struct {
	RoleID       uuid.UUID `json:"role_id" validate:"required"`
	PermissionID uuid.UUID `json:"permission_id" validate:"required"`
}

type UpdateRolePermission struct {
	ID           uuid.UUID  `json:"id" validate:"required"`
	RoleID       *uuid.UUID `json:"role_id" validate:"required"`
	PermissionID *uuid.UUID `json:"permission_id" validate:"required"`
}

type NewUserRole struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	RoleID uuid.UUID `json:"role_id" validate:"required"`
}

type UpdateUserRole struct {
	ID     uuid.UUID  `json:"id" validate:"required"`
	UserID *uuid.UUID `json:"user_id" validate:"required"`
	RoleID *uuid.UUID `json:"role_id" validate:"required"`
}
