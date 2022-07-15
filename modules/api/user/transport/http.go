package transport

import (
	"strconv"

	"bagogo/modules/api/user"
	"github.com/gin-gonic/gin"
)

// HTTP represents user http service
type HTTP struct {
	svc user.Service
}

func NewHTTP(svc user.Service, r *gin.RouterGroup) {
	h := HTTP{svc}

	ur := r.Group("/users")
	ur.GET("/:id", h.view)
}

func (h HTTP) view(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	result, err := h.svc.View(c, id)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}
