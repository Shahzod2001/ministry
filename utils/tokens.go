package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"ministry/config"
	"ministry/internal/model"
	"strings"
	"time"
)

var secretKey = []byte("ministry")

const (
	accessType  = "access"
	refreshType = "refresh"
)

func GenerateToken(entity interface{}) (*model.TokenPair, error) {
	var (
		id   int
		role string
	)

	switch e := entity.(type) {
	case *model.University:
		id = e.ID
		role = "university"
	case *model.Admin:
		id = e.ID
		role = "admin"
	}

	accessTokenTTL, err := time.ParseDuration(config.AppParams.TokenTTL.AccessTokenTTL)
	if err != nil {
		return nil, err
	}

	refreshTokenTTL, err := time.ParseDuration(config.AppParams.TokenTTL.RefreshTokenTTL)
	if err != nil {
		return nil, err
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         id,
		"role":       role,
		"token_type": accessType,
		"expiration": time.Now().Add(accessTokenTTL).Unix(),
	})

	signedAccess, err := accessToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         id,
		"token_type": refreshType,
		"expiration": time.Now().Add(refreshTokenTTL).Unix(),
	})

	signedRefresh, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &model.TokenPair{
		AccessToken:  signedAccess,
		RefreshToken: signedRefresh,
	}, nil

}

func GenerateTokenPairFromRefresh(refreshTokenString string) (*model.TokenPair, error) {
	claims, err := ParseToken(refreshTokenString)
	if err != nil {
		return nil, fmt.Errorf("refresh token: %w", err)
	}

	tokenType := claims["token_type"].(string)
	if tokenType != refreshType {
		return nil, fmt.Errorf("refresh token: Invalid token type")
	}

	refreshExp := time.Unix(int64(claims["expiration"].(float64)), 0)
	currentTime := time.Now()
	if currentTime.After(refreshExp) {
		return nil, errors.New("refresh token: Token is expired. Please login")
	}

	univerId, ok := claims["id"].(float64)
	if !ok {
		return nil, errors.New("invalid token claims: UniverId")
	}

	univer := model.University{
		ID: int(univerId),
	}

	return GenerateToken(&univer)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	if len(tokenString) == 0 {
		return nil, errors.New("empty token")
	}
	if len(strings.Split(tokenString, ".")) != 3 {
		return nil, errors.New("invalid jwt-token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
