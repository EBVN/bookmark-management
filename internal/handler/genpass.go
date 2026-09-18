package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/EBVN/bookmark-management/internal/service"
)

const (
	DefaultPasswordLength = 10
)

// GenPass interface
type GenPass interface {
	GeneratePassword(c *gin.Context)
}

type genPass struct {
	svc service.GenPass
}

// NewGenPass create new instance of GenPass
func NewGenPass(svc service.GenPass) GenPass {
	return &genPass{svc: svc}
}

// @Summary Generate a random password
// @Description Generates a cryptographically secure random password of fixed length.
// @Tags password
// @Produce application/json
// @Success 200 {object} map[string]string "message"
// @Failure 500 {object} string "Internal Server Error"
// @Router /genpass [get]
func (g *genPass) GeneratePassword(c *gin.Context) {
	// goi service
	pass, err := g.svc.GeneratePassword(DefaultPasswordLength)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	// tra ve response
	c.JSON(http.StatusOK, gin.H{"message": pass})
}
