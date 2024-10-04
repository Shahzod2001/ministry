package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func (h *Handler) NoRoute(c *gin.Context) {
	resp := response{
		Code:    http.StatusNotFound,
		Message: "Not Found",
	}
	c.JSON(http.StatusNotFound, gin.H{"error": resp})
}
