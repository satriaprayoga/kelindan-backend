package utils

import "github.com/golang-jwt/jwt"

type Claims struct {
	UserID   string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	UserType string `json:"user_type,omitempty"`
	jwt.StandardClaims
}
