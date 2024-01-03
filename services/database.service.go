package services

import (
	"example.com/mod/models"
	"example.com/mod/repository"
)

type UserService interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (*models.User, error)
	Create(user models.User) error
	Update(id int, user models.User) error
	Delete(id int) error
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

// Create implements UserService.
func (*UserServiceImpl) Create(user models.User) error {
	panic("unimplemented")
}

// Delete implements UserService.
func (*UserServiceImpl) Delete(id int) error {
	panic("unimplemented")
}

// GetByID implements UserService.
func (*UserServiceImpl) GetByID(id int) (*models.User, error) {
	panic("unimplemented")
}

// Update implements UserService.
func (*UserServiceImpl) Update(id int, user models.User) error {
	panic("unimplemented")
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

func (s *UserServiceImpl) GetAll() ([]models.User, error) {
	return s.userRepo.GetAll()
}
