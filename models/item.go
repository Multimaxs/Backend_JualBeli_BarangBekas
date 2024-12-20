package models

import "database/sql"

// Item struct for representing an item
type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"` // available, sold, etc.
	UserID      int     `json:"user_id"`
}

// AddItem for adding a new item
func AddItem(db *sql.DB, item Item) (int, error) {
	result, err := db.Exec(`
		INSERT INTO items (name, description, price, status, user_id)
		VALUES (?, ?, ?, ?, ?)
	`, item.Name, item.Description, item.Price, item.Status, item.UserID)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), nil
}
