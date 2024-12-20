package models

import "database/sql"

// User struct for representing an order
type Order struct {
	ID     int     `json:"id"`
	ItemID int     `json:"item_id"`
	UserID int     `json:"user_id"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

// CreateOrder for creating a new order
func CreateOrder(db *sql.DB, order Order) (int, error) {
	result, err := db.Exec(`
		INSERT INTO orders (item_id, user_id, status, amount)
		VALUES (?, ?, ?, ?)
	`, order.ItemID, order.UserID, order.Status, order.Amount)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), nil
}
