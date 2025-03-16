package handlers

import (
	"net/http"

	"github.com/bharabhi01/shorturl-go/internal/service"
	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	urlService *service.URLService
}

func NewURLHandler(urlService *service.URLService) *URLHandler {
	return &URLHandler{
		urlService: urlService,
	}
}

func (h *URLHandler) CreateShortURL(c *gin.Context) {
	var request struct {
		LongURL string `json:"long_url" binding:"required,url"`
		UserID  string `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL, err := h.urlService.CreateShortURL(c.Request.Context(), request.LongURL, request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_url": shortURL,
		"long_url":  request.LongURL,
	})
}

func (h *URLHandler) RedirectToLongURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	longURL, err := h.urlService.GetLongURL(c.Request.Context(), shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, longURL)
}
