package role

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

func (s *Store) CreateRole(role *types.NewRole) (*model.Role, error) {
	roleModel := &model.Role{}
	roleModel.Name = role.Name
	if role.Description != nil {
		roleModel.Description = role.Description
	}

	if err := s.db.Create(roleModel).Error; err != nil {
		return nil, err
	}
	return roleModel, nil
}

func (s *Store) GetRole(id uuid.UUID) (*model.Role, error) {
	var role model.Role
	if err := s.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *Store) GetRoles() ([]model.Role, error) {
	var roles []model.Role
	if err := s.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *Store) UpdateRole(role *types.UpdateRole) error {
	roleModel := &model.Role{}
	if err := s.db.First(roleModel, role.ID).Error; err != nil {
		return err
	}

	if role.Name != nil {
		roleModel.Name = *role.Name
	}

	if role.Description != nil {
		roleModel.Description = role.Description
	}

	if err := s.db.Save(roleModel).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteRole(id uuid.UUID) error {
	roleModel := &model.Role{}
	if err := s.db.First(roleModel, id).Error; err != nil {
		return err
	}

	if err := s.db.Delete(roleModel).Error; err != nil {
		return err
	}
	return nil
}
