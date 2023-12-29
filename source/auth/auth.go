package auth

import (
	"authentication/source/model"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HashPwd(password string) string {
	passSha256 := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(passSha256[:])
}

func GenerateAccessToken(ctx context.Context, id primitive.ObjectID, tokenExpired int, jwtSecret string) (
	string, error) {

	return jwt.NewWithClaims(jwt.SigningMethodHS256, model.UserClaim{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, tokenExpired)),
		},
	}).SignedString([]byte(jwtSecret))

}

func GenerateRefreshToken(ctx context.Context, id primitive.ObjectID, tokenExpired int, jwtSecret string) (
	string, error) {

	return jwt.NewWithClaims(jwt.SigningMethodHS256, model.UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, tokenExpired)),
		},
	}).SignedString([]byte(jwtSecret))

}

func ValidateAT(token string, jwtSecret string) (id primitive.ObjectID, err error) {
	parsedToken, err := jwt.ParseWithClaims(token, &model.UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		println("Validate signature error:", err.Error())
	}
	if parsedToken == nil {
		return primitive.NilObjectID, errors.New("can not parse token")
	}

	claims, ok := parsedToken.Claims.(*model.UserClaim)
	if !ok {
		return primitive.NilObjectID, errors.New("can not parse token")
	}
	dt := claims.ExpiresAt.Time.Unix() - time.Now().Unix()
	if dt < 0 {
		return primitive.NilObjectID, errors.New("jwt expired")
	}
	return claims.Id, nil
}
