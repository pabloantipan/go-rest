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

func NewUserRepository() *UserRepository {
	return &UserRepository{db: PgConnect()}
}

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

func PgConnect() *sql.DB {
	psqlInfo := "host=localhost port=5432 user=myuser password=mypassword dbname=mydatabase sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

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

func (repo *UserRepository) DropUserTable() {
	query := "DROP TABLE users IF EXISTS"
	result, err := repo.db.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("DropUserTable()", result)
}
