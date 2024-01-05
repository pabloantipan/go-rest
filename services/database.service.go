package services

import (
	"database/sql"

	"example.com/mod/models"
	"example.com/mod/repository"
)

type UserDBService interface {
	// NewDBUserService() *UserDBService
	GetAll() ([]models.User, error)
	GetByID(id int) (*models.User, error)
	Create(user models.User) (sql.Result, error)
	Update(id int, user models.User) error
	Delete(id int) error
	DropUserTable()
}

type UserDBServiceImpl struct {
	userRepo repository.UserRepository
}

func NewDBUserService(userRepo repository.UserRepository) UserDBService {
	return &UserDBServiceImpl{userRepo: userRepo}
}

func (u *UserDBServiceImpl) GetAll() ([]models.User, error) {
	rows, err := u.userRepo.RunQuery("SELECT * FROM users")

	if err != nil {
		panic(err)
	}

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Create implements UserDBService.
func (u *UserDBServiceImpl) Create(user models.User) (sql.Result, error) {
	return u.userRepo.Exec(user)
	// result, err := db.Exec("INSERT INTO users (name) VALUES ($1)", "John Doe")
}

// Delete implements UserDBService.
func (*UserDBServiceImpl) Delete(id int) error {
	panic("unimplemented")
}

// GetByID implements UserDBService.
func (*UserDBServiceImpl) GetByID(id int) (*models.User, error) {
	panic("unimplemented")
}

// Update implements UserDBService.
func (*UserDBServiceImpl) Update(id int, user models.User) error {
	panic("unimplemented")
}

func (u *UserDBServiceImpl) DropUserTable() {
	u.userRepo.DropUserTable()
}
