package models

import "database/sql"

// UserDao makes the commom database operations
type UserDao struct {
	Conn *sql.DB
}

// NewUserDao
func NewUserDao(db *sql.DB) *UserDao {
	return &UserDao{
		Conn: db,
	}
}

// Create a new user instance
func (userDao *UserDao) Create(user User) error {
	stmt, err := userDao.Conn.Prepare("INSERT INTO users(name, email, password) VALUES($1, $2, $3)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Name, user.Email, user.Password)
	return err
}

// Find retrieves the user with correct email and password
func (userDao *UserDao) Find(email string, password string) (*User, error) {
	var user User
	row := userDao.Conn.QueryRow("SELECT * FROM users WHERE email=$1 AND password=$2", email, password)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
