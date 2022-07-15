package user

import (
	"bagogo/modules/entity"
	"github.com/gin-gonic/gin"
)

// View returns single user
func (u User) View(c *gin.Context, id int) (entity.User, error) {
	return u.dbe.View(u.db, id)
}