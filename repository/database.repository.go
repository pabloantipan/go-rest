package repository

import (
	"database/sql"
	"fmt"

	"example.com/mod/models"
	_ "github.com/lib/pq"
)

// type UserRepository interface {
// 	NewUserRepository() *UserRepository
// 	runQuery
// }

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: Connect()}
}

// func (r *UserRepository) GetAll() ([]models.User, error) {
// 	rows, err := r.db.Query("SELECT * FROM users")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	users := []models.User{}
// 	for rows.Next() {
// 		var user models.User
// 		err := rows.Scan(&user.ID, &user.Name, &user.Email)
// 		if err != nil {
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}

// 	return users, nil
// }

func (repo *UserRepository) RunQuery(query string) (*sql.Rows, error) {
	return repo.db.Query(query)
}

func (repo *UserRepository) Exec(user models.User) (sql.Result, error) {
	return repo.db.Exec(
		"INSERT INTO users (name, email) VALUES ($1, $2)",
		user.Name,
		user.Email,
	)
}

func Connect() *sql.DB {
	psqlInfo := "host=localhost port=5432 user=myuser password=mypassword dbname=mydatabase sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	CreateUsersTable(db)

	rows, err := db.Query("SELECT * FROM users")
	fmt.Println("rows", rows)

	return db
}

func CreateUsersTable(db *sql.DB) {
	query := "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)"
	result, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("CreateUsersTable()", result)
}
