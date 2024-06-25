package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// UserService interface
type UserService interface {
	Register(input CreateUserInput) (UserTable, error)
	UpdateUser(id int, input UpdateUserInput) (UserTable, error)
	Login(input LoginInput) (UserTable, error)
	GetUserByID(id int) (UserTable, error)
	GetUsers() ([]UserTable, error)
}

// userService struct
type userService struct {
	repository UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repository UserRepository) *userService {
	return &userService{repository}
}

// Register registers a new user
func (s *userService) Register(input CreateUserInput) (UserTable, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserTable{}, err
	}
	_, err = s.repository.FindByUsername(input.Username)
	if err == nil {
		return UserTable{}, errors.New("username already taken")
	}


	user := UserTable{
		Nama:     input.Nama,
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	newUser, err := s.repository.Save(user)
	return newUser, err
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(id int, input UpdateUserInput) (UserTable, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return UserTable{}, err
	}
	if user.Username != input.Username {
		_, err := s.repository.FindByUsername(input.Username)
		if err == nil {
			return UserTable{}, errors.New("username already taken")
		}
	}

	user.Nama = input.Nama
	user.Username = input.Username
	user.Role = input.Role

	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return UserTable{}, err
		}
		user.Password = string(hashedPassword)
	}

	updatedUser, err := s.repository.Update(user)
	return updatedUser, err
}

// Login authenticates a user
func (s *userService) Login(input LoginInput) (UserTable, error) {
	user, err := s.repository.FindByUsername(input.Username)
	if err != nil {
		return UserTable{}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return UserTable{}, errors.New("invalid credentials")
	}

	return user, nil
}

// GetUserByID retrieves a user by its ID
func (s *userService) GetUserByID(id int) (UserTable, error) {
	user, err := s.repository.FindByID(id)
	return user, err
}

// GetUsers retrieves all users
func (s *userService) GetUsers() ([]UserTable, error) {
	users, err := s.repository.FindAll()
	return users, err
}
