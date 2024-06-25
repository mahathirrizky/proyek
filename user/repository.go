package user

import (
	"gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
	Save(user UserTable) (UserTable, error)
	FindByID(id int) (UserTable, error)
	FindAll() ([]UserTable, error)
	FindByUsername(username string) (UserTable, error)
	Update(user UserTable) (UserTable, error)
}

// userRepository struct
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

// Save inserts a new user record in the database
func (r *userRepository) Save(user UserTable) (UserTable, error) {
	err := r.db.Create(&user).Error
	return user, err
}

// FindByID retrieves a user record by its ID
func (r *userRepository) FindByID(id int) (UserTable, error) {
	var user UserTable
	err := r.db.First(&user, id).Error
	return user, err
}

// FindAll retrieves all user records from the database
func (r *userRepository) FindAll() ([]UserTable, error) {
	var users []UserTable
	err := r.db.Find(&users).Error
	return users, err
}

// FindByUsername retrieves a user record by its username
func (r *userRepository) FindByUsername(username string) (UserTable, error) {
	var user UserTable
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

// Update updates an existing user record in the database
func (r *userRepository) Update(user UserTable) (UserTable, error) {
	err := r.db.Save(&user).Error
	return user, err
}
