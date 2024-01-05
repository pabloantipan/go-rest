package services

import (
	"example.com/mod/models"
	"example.com/mod/repository"
)

type UserCacheService interface {
	GetAll() ([]models.User, error)
	GetByID(id string) (models.User, error)
	Create(user models.User) models.CacheCreateResponse
	Update(id string, user models.User) error
	Delete(id string) error
}

type UserCacheServiceImpl struct {
	userCacheRepository repository.UserCacheRepository
}

func NewCacheUserService(userCacheRepository repository.UserCacheRepository) UserCacheService {
	return &UserCacheServiceImpl{userCacheRepository: userCacheRepository}
}

func (u *UserCacheServiceImpl) GetAll() ([]models.User, error) {
	return u.userCacheRepository.GetAll()
}

func (u *UserCacheServiceImpl) GetByID(cacheID string) (models.User, error) {
	return u.userCacheRepository.GetByID(cacheID)
}

func (u *UserCacheServiceImpl) Create(user models.User) models.CacheCreateResponse {
	return u.userCacheRepository.CreateUser(user)
}

func (u *UserCacheServiceImpl) Update(id string, user models.User) error {
	panic("unimplemented")
}

func (u *UserCacheServiceImpl) Delete(id string) error {
	panic("unimplemented")
}
