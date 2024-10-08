package user

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"player/internal/config"
)

type User struct {
	Username string `json:"userLogin"`
	Password string `json:"userPassword"`
	Email    string `json:"userEmail"`
	Remember bool   `json:"userRemember"`
}

func (u *User) AddUser(user User) error {
	cnf := config.DefaultConfig()

	connStr := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.DBConfig.Port, cnf.DBConfig.User, cnf.DBConfig.Password, cnf.DBConfig.DBName, cnf.DBConfig.SslMode)
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		return err
	}

	HashPassword, err := HashPassword(user.Password)

	if err != nil {
		return err
	}

	sqlRequest := `INSERT INTO users (Username, Email, HashPassword) VALUES ($1, $2, $3)`

	_, err = db.Exec(sqlRequest, user.Username, user.Email, HashPassword)

	if err != nil {
		return err
	}

	return nil
}

func HashPassword(pass string) (string, error) {
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	return string(HashPassword), err
}
