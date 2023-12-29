package model

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTResult struct {
	Id           primitive.ObjectID `json:"_id"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refreshToken"`
}
type UserClaim struct {
	Id primitive.ObjectID `json:"_id"`
	jwt.RegisteredClaims
}
