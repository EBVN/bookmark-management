// Package handler provides HTTP request handlers for the bookmark management service.
// It contains handlers for various endpoints including password generation and user management.
package handler

import (
	"net/http"

	"github.com/EBVN/bookmark-management/internal/service"
	"github.com/gin-gonic/gin"
)

type GenPass interface {
	GeneratePassword(c *gin.Context)
}

const DefaultLength = 10

type genPass struct {
	svc service.GenPass
}

func NewGenPass(svc service.GenPass) GenPass {
	return &genPass{svc: svc}
}

// GeneratePassword
// @Summary Generate a random password
// @Description Generates a cryptographically secure random password of fixed length.
// @Tags password
// @Produce application/json
// @Success 200 {object} map[string]string "message"
// @Failure 500 {object} string "Internal Server Error"
// @Router /genpass [get]
func (g *genPass) GeneratePassword(c *gin.Context) {
	// call service
	pass, err := g.svc.GeneratePassword(DefaultLength)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// return result
	c.JSON(http.StatusOK, gin.H{"password": pass})
}
