package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email        string             `json:"email" bson:"email,omitempty"`
	Password     string             `json:"password" bson:"password,omitempty"`
	Token        string             `json:"token" bson:"token,omitempty"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token,omitempty"`
	Created      time.Time          `json:"created" bson:"created"`
}

type RegisterUserInput struct {
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Code    int    `json:"code" `
	Message string `json:"message"`
	Result  Result `json:"result"`
}
type GetUserResponse struct {
	Code    int    `json:"code" `
	Message string `json:"message"`
	Result  User   `json:"result"`
}
type JobDescriptionResponse struct {
	Code    int            `json:"code" `
	Message string         `json:"message"`
	Result  JobDescription `json:"result"`
}

type JobDescription struct {
	Company  string `json:"company" bson:"company,omitempty"`
	Age      int    `json:"age" bson:"age,omitempty"`
	Position string `json:"position" bson:"position,omitempty"`
}

type Result struct {
	ID           primitive.ObjectID `json:"id,omitempty"`
	Token        string             `json:"token" bson:"token,omitempty"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token,omitempty"`
}

type CheckTokenReq struct {
	Token string `json:"token"`
}

type CheckTokenRespond struct {
	Code    int    `json:"code" `
	Message string `json:"message"`
}
type LogoutRespond struct {
	Code    int    `json:"code" `
	Message string `json:"message"`
}

// func FilteredResponse(user *User) User {
// 	return User{
// 		ID:      user.ID,
// 		Email:   user.Email,
// 		Created: user.Created,
// 	}
// }
func (r LoginUserInput) Validate() error {
	if len(r.Email) == 0 || len(r.Password) == 0 {
		return errors.New("username or password invalid")
	}

	return nil
}
func (r User) Validate() error {
	if len(r.Email) == 0 || len(r.Password) == 0 {
		return errors.New("username or password invalid")
	}

	return nil
}
