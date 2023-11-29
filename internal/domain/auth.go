package domain

import "github.com/golang-jwt/jwt"

// Token Claims
type StandardClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

type UserAccessTokenClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

type UserRefreshTokenClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

type RefreshTokenLog struct {
	ID              int64  `sql:"id,omitempty"`
	RefreshToken    string `sql:"refresh_token,omitempty"`
	UserID          int64  `sql:"user_id,omitempty"`
	UUID            string `sql:"uuid,omitempty"`
	Status          string `sql:"status,omitempty"`
	LastUsedAtStamp int64  `sql:"last_used_at,omitempty"`
	CreatedAtStamp  int64  `sql:"created_at_stamp,omitempty"`
	UpdatedAtStamp  int64  `sql:"updated_at_stamp,omitempty"`
}
