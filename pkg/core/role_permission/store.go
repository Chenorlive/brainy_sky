package rolepermission

import (
	"gorm.io/gorm"

	"github.com/Chenorlive/brainy/model"
	"github.com/Chenorlive/brainy/types"
	"github.com/gofrs/uuid"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateRolePermission(rolePermission *types.NewRolePermission) (*model.RolePermission, error) {
	rolePermissionModel := &model.RolePermission{
		RoleID:       rolePermission.RoleID,
		PermissionID: rolePermission.PermissionID,
	}

	if err := s.db.Create(rolePermissionModel).Error; err != nil {
		return nil, err
	}
	return rolePermissionModel, nil
}

func (s *Store) GetRolePermission(id uuid.UUID) (*model.RolePermission, error) {
	var rolePermission model.RolePermission
	if err := s.db.Preload("Role").Preload("Permission").First(&rolePermission, id).Error; err != nil {
		return nil, err
	}
	return &rolePermission, nil
}

func (s *Store) GetRolePermissions() ([]model.RolePermission, error) {
	var rolePermissions []model.RolePermission
	if err := s.db.Preload("Role").Preload("Permission").Find(&rolePermissions).Error; err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (s *Store) UpdateRolePermission(rolePermission *types.UpdateRolePermission) error {
	rolePermissionModel := &model.RolePermission{}
	if err := s.db.First(rolePermissionModel, rolePermission.ID).Error; err != nil {
		return err
	}

	if rolePermission.RoleID != nil {
		rolePermissionModel.RoleID = *rolePermission.RoleID
	}
	if rolePermission.PermissionID != nil {
		rolePermissionModel.PermissionID = *rolePermission.PermissionID
	}

	if err := s.db.Save(rolePermissionModel).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteRolePermission(id uuid.UUID) error {
	rolePermissionModel := &model.RolePermission{}
	if err := s.db.First(rolePermissionModel, id).Error; err != nil {
		return err
	}

	if err := s.db.Delete(rolePermissionModel).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) GetRolePermissionsByRoleID(roleID uuid.UUID) ([]model.RolePermission, error) {
	var rolePermissions []model.RolePermission
	if err := s.db.Preload("Role").Preload("Permission").Where("role_id = ?", roleID).Find(&rolePermissions).Error; err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (s *Store) GetRolePermissionsByPermissionID(permissionID uuid.UUID) ([]model.RolePermission, error) {
	var rolePermissions []model.RolePermission
	if err := s.db.Preload("Role").Preload("Permission").Where("permission_id = ?", permissionID).Find(&rolePermissions).Error; err != nil {
		return nil, err
	}
	return rolePermissions, nil
}
