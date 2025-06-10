package user

import (
	"fmt"

	"github.com/Chenorlive/brainy/model"
	"github.com/Chenorlive/brainy/pkg/auth"
	userrole "github.com/Chenorlive/brainy/pkg/core/userRole"
	"github.com/Chenorlive/brainy/types"
	"gorm.io/gorm"
)

type Store struct {
	db            *gorm.DB
	userRoleStore *userrole.Store
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db:            db,
		userRoleStore: userrole.NewStore(db),
	}
}

func (s *Store) CreateUser(user *types.RegisterUser) (*model.User, error) {
	password, err := auth.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	var role model.Role
	if err := s.db.First(&role, "name = ?", user.Role).Error; err != nil {
		return nil, fmt.Errorf("role not found: %w", err)
	}

	userModel := &model.User{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  password,
		LoginHint: &user.Password, // Password should be hashed before storing
	}

	if user.MiddleName != nil {
		userModel.MiddleName = user.MiddleName
	}

	if user.Email != nil {
		userModel.Email = user.Email
	}

	if user.Phone != nil {
		userModel.Phone = user.Phone
	}
	if user.DateOfBirth != nil {
		userModel.DateOfBirth = user.DateOfBirth
	}

	if user.Address != nil {
		userModel.Address = user.Address
	}

	if user.Image != nil {
		userModel.Image = user.Image
	}

	tx := s.db.Begin()

	// Create the user
	if err := tx.Create(userModel).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// assign the role to the user
	userRole := &model.UserRole{
		UserID: userModel.ID,
		RoleID: role.ID,
	}

	if err := tx.Create(userRole).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return userModel, nil
}

func (s *Store) GetUser(id uint) (*model.User, error) {
	var user model.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Store) GetUsers() ([]model.User, error) {
	var users []model.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Store) UpdateUser(user *types.UpdateUser) error {
	userModel := &model.User{}
	if err := s.db.First(userModel, user.ID).Error; err != nil {
		return err
	}

	if user.Username != nil {
		userModel.Username = *user.Username
	}
	if user.FirstName != nil {
		userModel.FirstName = *user.FirstName
	}
	if user.MiddleName != nil {
		userModel.MiddleName = user.MiddleName
	}
	if user.LastName != nil {
		userModel.LastName = *user.LastName
	}
	if user.Email != nil {
		userModel.Email = user.Email
	}
	if user.Phone != nil {
		userModel.Phone = user.Phone
	}
	if user.DateOfBirth != nil {
		userModel.DateOfBirth = user.DateOfBirth
	}
	if user.Address != nil {
		userModel.Address = user.Address
	}
	if user.Image != nil {
		userModel.Image = user.Image
	}

	if err := s.db.Save(userModel).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteUser(id uint) error {
	userModel := &model.User{}
	if err := s.db.First(userModel, id).Error; err != nil {
		return err
	}

	if err := s.db.Delete(userModel).Error; err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Store) UpdatePassword(user types.UpdatePassword) (*model.User, error) {
	userModel := &model.User{}
	if err := s.db.First(userModel, user.ID).Error; err != nil {
		return nil, err
	}

	ok := auth.CheckPasswordHash(user.OldPassword, userModel.Password)
	if ok != true {
		return nil, fmt.Errorf("old password is incorrect")
	}

	hashedPassword, err := auth.HashPassword(user.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to hash new password: %w", err)
	}
	userModel.Password = hashedPassword
	userModel.LoginHint = &user.NewPassword // Update login hint with the new password

	if err := s.db.Save(userModel).Error; err != nil {
		return nil, err
	}
	return userModel, nil
}
