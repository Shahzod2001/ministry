package handler

import (
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, resp *response, statusCode int, err error) {
	resp.Code = statusCode
	resp.Message = err.Error()
	c.JSON(statusCode, gin.H{"error": resp})
}
