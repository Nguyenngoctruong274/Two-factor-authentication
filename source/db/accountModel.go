package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"uuid" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email,omitempty"`
	Password     string             `json:"password" bson:"password,omitempty"`
	Token        string             `json:"token" bson:"token,omitempty"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token,omitempty"`
	Created      time.Time          `json:"created" bson:"created"`
}

type JobDecription struct {
	Company  string `json:"company" bson:"company,omitempty"`
	Age      int    `json:"age" bson:"age,omitempty"`
	Position string `json:"position" bson:"position,omitempty"`
}

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID           string    `json:"id,omitempty"`
	Email        string    `json:"email" bson:"email,omitempty"`
	Password     string    `json:"password" bson:"password,omitempty"`
	Token        string    `json:"token" bson:"token,omitempty"`
	RefreshToken string    `json:"refresh_token" bson:"refresh_token,omitempty"`
	Created      time.Time `json:"created" bson:"created"`
}

func FilteredResponse(user *User) UserResponse {
	return UserResponse{
		ID:      user.ID.String(),
		Email:   user.Email,
		Created: user.Created,
	}
}
