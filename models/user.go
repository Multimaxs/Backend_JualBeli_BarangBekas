package models

import "database/sql"

// User struct for representing a user
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // buyer or seller
}

// CreateOrder for creating a new user
func CreateUser(db *sql.DB, user User) (int, error) {
	result, err := db.Exec(`
		INSERT INTO users (name, email, password, role)
		VALUES (?, ?, ?, ?)
	`, user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), nil
}
