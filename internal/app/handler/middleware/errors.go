package middleware

import (
	"errors"
)

var (
	errAuthorizationMissing          = errors.New("authorization header is missing")
	errInvalidTokenFormat            = errors.New("invalid token format")
	errIncorrectTokenType            = errors.New("incorrect token type")
	errTokenExpirationClaimsNotFound = errors.New("token expiration claims not found")
	errAccessTokenExpired            = errors.New("access token expired")
)
