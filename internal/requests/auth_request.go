package requests

import "github.com/golang-jwt/jwt/v4"

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaims struct {
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
	jwt.RegisteredClaims
}
