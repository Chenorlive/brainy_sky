package userrole

import (
	"fmt"
	"log"

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

func (s *Store) DeleteUserRolesByUserID(userID uuid.UUID) error {
	var userRoles []*model.UserRole
	if err := s.db.Where("user_id = ?", userID).Find(&userRoles).Error; err != nil {
		return err
	}

	for _, userRole := range userRoles {
		if err := s.db.Delete(userRole).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) GetUserRolesByUserID(userID uuid.UUID) ([]*model.UserRole, error) {
	var userRoles []*model.UserRole
	if err := s.db.Preload("User").Preload("Role").Where("user_id = ?", userID).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (s *Store) GetUserPermissions(userID uuid.UUID) []*model.Permission {
	var permissions []*model.Permission

	err := s.db.
		Model(&model.Permission{}).
		Distinct("permissions.id").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Find(&permissions).Error

	if err != nil {
		log.Fatalf("Error fetching permissions: %v", err)
	}

	fmt.Printf("Permissions for User %s:\n", userID)
	for _, perm := range permissions {
		fmt.Printf("- Name: %s, ID: %s\n", perm.Name, perm.ID)
	}

	return permissions
}
