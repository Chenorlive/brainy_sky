package userrole

import (
	"github.com/Chenorlive/brainy/model"
	"github.com/Chenorlive/brainy/types"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateUserRole(userRole *types.NewUserRole) (*model.UserRole, error) {
	userRoleModel := &model.UserRole{
		UserID: userRole.UserID,
		RoleID: userRole.RoleID,
	}

	if err := s.db.Create(userRoleModel).Error; err != nil {
		return nil, err
	}
	return userRoleModel, nil
}

func (s *Store) GetUserRole(id uuid.UUID) (*model.UserRole, error) {
	var userRole model.UserRole
	if err := s.db.Preload("User").Preload("Role").First(&userRole, id).Error; err != nil {
		return nil, err
	}
	return &userRole, nil
}

func (s *Store) GetUserRoles() ([]*model.UserRole, error) {
	var userRoles []*model.UserRole
	if err := s.db.Preload("User").Preload("Role").Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (s *Store) UpdateUserRole(userRole *types.UpdateUserRole) error {
	userRoleModel := &model.UserRole{}
	if err := s.db.First(userRoleModel, userRole.ID).Error; err != nil {
		return err
	}

	if userRole.UserID != nil {
		userRoleModel.UserID = *userRole.UserID
	}
	if userRole.RoleID != nil {
		userRoleModel.RoleID = *userRole.RoleID
	}

	if err := s.db.Save(userRoleModel).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteUserRole(id uuid.UUID) error {
	userRoleModel := &model.UserRole{}
	if err := s.db.First(userRoleModel, id).Error; err != nil {
		return err
	}

	if err := s.db.Delete(userRoleModel).Error; err != nil {
		return err
	}
	return nil
}
