package repository

import (
	"bagogo/modules/entity"
	"github.com/jmoiron/sqlx"
)

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u User) View(db *sqlx.DB, id int) (entity.User, error) {
	var user entity.User

	err := db.Get(&user, "SELECT id, name, email, phone FROM users where id = ?", id)
	return user, err
}
