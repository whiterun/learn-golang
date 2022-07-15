package user

import (
	"bagogo/modules/entity"
	"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
	"bagogo/modules/api/user/repository"
)

// User represents user application service
type User struct {
	db *sqlx.DB
	dbe DBE
}

// DBE represents user repository interface
type DBE interface {
	View(*sqlx.DB, int) (entity.User, error)
}

// Service represents user application interface
type Service interface {
	View(*gin.Context, int) (entity.User, error)
}

// New creates new user application service
func New(db *sqlx.DB, dbe DBE) *User {
	return &User{db: db, dbe: dbe}
}

// Initialize initalizes User application service with defaults
func Initialize(db *sqlx.DB) *User {
	return New(db, repository.User{})
}
