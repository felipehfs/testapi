package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// UserDao makes the commom database operations
type UserDao struct {
	Conn *sql.DB
}

// NewUserDao instantiate userDado
func NewUserDao(db *sql.DB) *UserDao {
	return &UserDao{
		Conn: db,
	}
}

// Create a new user instance
func (userDao *UserDao) Create(user User) (*User, error) {
	query := `
		INSERT INTO 
			users(name, email, password) 
		VALUES($1, $2, $3)
		RETURNING id
	`
	err := userDao.Conn.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	return &user, err
}

// Find retrieves the user with correct email and password
func (userDao *UserDao) Find(email, password string) (*User, error) {
	var user User
	row := userDao.Conn.QueryRow("SELECT * FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return &user, nil
}
