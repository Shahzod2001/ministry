package middleware

import (
	"github.com/gin-gonic/gin"
	"ministry/utils"
	"net/http"
	"strings"
	"time"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func abortWithError(c *gin.Context, resp *response, statusCode int, err error) {
	resp.Code = statusCode
	resp.Message = err.Error()
	c.JSON(statusCode, gin.H{"error": resp})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, OPTIONS, PUT, PATCH, HEAD")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

const (
	accessType  = "access"
	refreshType = "refresh"
)

func AuthMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp response
		claims, err := GetClaimsInAccessToken(c)
		if err != nil {
			resp.Code = http.StatusUnauthorized
			resp.Message = err.Error()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": resp})
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			resp.Code = http.StatusUnauthorized
			resp.Message = "invalid token claims: Role"
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": resp})
			return
		}

		univerId, ok := claims["id"].(float64)
		if !ok {
			resp.Code = http.StatusUnauthorized
			resp.Message = "invalid token claims: UniverId"
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": resp})
			return
		}

		c.Set("role", role)
		c.Set("univer_id", int(univerId))

		c.Next()
	}
}

func AdminMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp response
		role := c.GetString("role")

		if role != "admin" {
			resp.Code = http.StatusForbidden
			resp.Message = "Access denied"
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": resp})
			return
		}

		c.Next()
	}
}

func GetClaimsInAccessToken(c *gin.Context) (map[string]interface{}, error) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return nil, errAuthorizationMissing
	}

	arrToken := strings.Split(accessToken, " ")
	if len(arrToken) != 2 {
		return nil, errInvalidTokenFormat
	}

	tokenString := arrToken[1]

	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	tokenType, ok := claims["token_type"].(string)
	if !ok || tokenType != accessType {
		return nil, errIncorrectTokenType
	}

	expiration, ok := claims["expiration"].(float64)
	if !ok {
		return nil, errTokenExpirationClaimsNotFound
	}

	accessExp := time.Unix(int64(expiration), 0)
	currentTime := time.Now()
	if currentTime.After(accessExp) {
		return nil, errAccessTokenExpired
	}

	return claims, nil
}
