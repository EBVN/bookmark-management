package handler

import (
	"errors"
	"net/http"

	"github.com/EBVN/bookmark-management/internal/service"
	"github.com/gin-gonic/gin"
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
	Exp int    `json:"exp" binding:"gte=0,lte=60000"`
	URL string `json:"url" binding:"required,url"`
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
func (h *shortenLink) CreateShortenLink(c *gin.Context) {
	// lay input
	input := &shortenLinkInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	// goi svc
	code, err := h.svc.CreateShortenLink(c, input.URL)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// return resp
	c.JSON(http.StatusOK, gin.H{"code": code})
}

// Redirect Forward the request to the original url
// @Tags link
// @Accept application/json
// @Produce application/json
// @Param code path string true "Shorten code"
// @Success 302
// @Router /v1/links/redirect/{code} [get]
func (h *shortenLink) Redirect(c *gin.Context) {
	// lay input
	code := c.Param("code")
	if code == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	// goi svc
	url, err := h.svc.GetUrlFromCode(c, code)
	if err != nil {
		if errors.Is(err, service.ErrUrlDoesNotExist) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Code Not Exist"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)

}
