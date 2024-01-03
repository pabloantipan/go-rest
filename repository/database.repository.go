package repository

import (
	"database/sql"
	"fmt"

	"example.com/mod/models"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func Connect() *sql.DB {
	psqlInfo := "host=localhost port=5432 user=your_username password=your_password dbname=your_database sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM users")
	fmt.Println("rows", rows)

	result, err := db.Exec("INSERT INTO users (name) VALUES ($1)", "John Doe")
	fmt.Println("result", result)

	// db.Close()
	return db
}
