package permission

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

func (s *Store) CreatePermission(permission *types.NewPermission) (*model.Permission, error) {
	permissionModel := &model.Permission{
		Name:        permission.Name,
		Description: permission.Description,
	}

	if err := s.db.Create(permissionModel).Error; err != nil {
		return nil, err
	}
	return permissionModel, nil
}

func (s *Store) GetPermission(id uuid.UUID) (*model.Permission, error) {
	permission := &model.Permission{}
	if err := s.db.First(permission, id).Error; err != nil {
		return nil, err
	}
	return permission, nil
}

func (s *Store) GetPermissions() ([]*model.Permission, error) {
	permissions := []*model.Permission{}
	if err := s.db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (s *Store) UpdatePermission(permission *types.UpdatePermission) error {
	permissionModel := &model.Permission{}
	if err := s.db.First(permissionModel, permission.ID).Error; err != nil {
		return err
	}

	if permission.Name != nil {
		permissionModel.Name = *permission.Name
	}
	if permission.Description != nil {
		permissionModel.Description = permission.Description
	}

	if err := s.db.Save(permissionModel).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) DeletePermission(id uuid.UUID) error {
	permissionModel := &model.Permission{}
	if err := s.db.First(permissionModel, id).Error; err != nil {
		return err
	}

	if err := s.db.Delete(permissionModel).Error; err != nil {
		return err
	}
	return nil
}
