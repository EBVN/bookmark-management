package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/EBVN/bookmark-management/internal/repository"
	"github.com/EBVN/bookmark-management/internal/service"
)

type ShortenLink interface {
	CreateShortenLink(c *gin.Context)
	Redirect(c *gin.Context)
}

type shortenLink struct {
	svc service.ShortenUrl
}

func NewShortenLink(svc service.ShortenUrl) ShortenLink {
	return &shortenLink{svc: svc}
}

type shortenLinkInput struct {
	Url string `json:"url" binding:"required,url"`
	Exp int64  `json:"exp" binding:"required,gte=500"`
}

// CreateShortenLink   Generate shorten link
// @Summary      Generate shorten url based on original url that last upto 7 days
// @Description  Generate shorten url based on original url that last upto 7 days
// @Tags         link
// @Accept       application/json
// @Produce      application/json
// @Param        input body shortenLinkInput true "Input required"
// @Success      200 {object} map[string]string
// @Router       /v1/links/shorten [post]
func (s *shortenLink) CreateShortenLink(c *gin.Context) {
	// lay input
	input := &shortenLinkInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// goi svc
	key, err := s.svc.CreateShortenLink(c, input.Url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create shorten url - CreateShortenLink endpoint")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// tra ve response
	c.JSON(http.StatusOK, gin.H{"key": key})
}

// Redirect Forward the request to the original url
// @Tags link
// @Accept application/json
// @Produce application/json
// @Param code path string true "Shorten code"
// @Success 302
// @Router /v1/links/redirect/{code} [get]
func (s *shortenLink) Redirect(c *gin.Context) {
	code := c.Param("code")
	log.Debug().Msgf("code: %s", code)
	if code == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid code"})
		return
	}

	url, err := s.svc.GetUrlFromCode(c, code)
	if err != nil {
		if errors.Is(err, repository.ErrKeyNotFound) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
			return
		}

		log.Error().Err(err).Msg("Failed to get url from code - Redirect endpoint")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
