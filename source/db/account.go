package db

import (
	"authentication/source/auth"
	"authentication/source/model"
	"context"
	"os"
	"strconv"
	"time"

	"github.com/ioVN/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountDB struct {
	ac *mongo.Collection
}

func (a *AccountDB) GenerateToken(ctx context.Context, id primitive.ObjectID, username, password string) (
	*model.JWTResult, error) {
	//Gen JWT
	tokenExpiredStr := os.Getenv("TOKEN_MAXAGE")
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenExpired, _ := strconv.Atoi(tokenExpiredStr)

	token, err := auth.GenerateAccessToken(ctx, id, (tokenExpired), jwtSecret)
	if err != nil {
		return nil, err
	}

	refresh, err := auth.GenerateRefreshToken(ctx, id, (tokenExpired), jwtSecret)
	if err != nil {
		return nil, err
	}

	return &model.JWTResult{
		Id:           id,
		Token:        token,
		RefreshToken: refresh,
	}, nil

}

func (a *AccountDB) CreateUser(ctx context.Context, email, password string) error {
	id := primitive.NewObjectID()

	if _, err := a.ac.InsertOne(ctx, &User{
		ID:       id,
		Email:    email,
		Password: password,
		Created:  time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

func (s *AccountDB) UpdateJWT(ctx context.Context, id primitive.ObjectID, jwt string) error {
	var updatedDocument bson.M
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"token": jwt}}
	if err := s.ac.FindOneAndUpdate(ctx, filter, update, nil).Decode(&updatedDocument); err != nil {
		return err
	}

	return nil

}
func (a *AccountDB) GetAccountById(ctx context.Context, id string) (*model.User, error) {

	u := &model.User{}
	idObj, _ := primitive.ObjectIDFromHex(id)
	res := a.ac.FindOne(ctx, bson.M{
		"_id": idObj,
	}, nil)

	if err := res.Decode(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (a *AccountDB) GetAccountByToken(ctx context.Context, token string) (*model.User, error) {

	u := &model.User{}

	res := a.ac.FindOne(ctx, bson.M{
		"token": token,
	}, nil)

	if err := res.Decode(u); err != nil {
		return nil, err
	}

	return u, nil
}
func (a *AccountDB) GetAccount(ctx context.Context, email, pwd string) (*model.User, error) {

	u := &model.User{}

	res := a.ac.FindOne(ctx, bson.M{
		"email":    email,
		"password": pwd,
	}, nil)

	if err := res.Decode(u); err != nil {
		return nil, err
	}

	return u, nil
}

func NewAccoutDB(db *mongo.Database) *AccountDB {
	return &AccountDB{
		ac: database.MongoInit(
			db,
			"account",
			database.MakeMongoIndex(bson.D{
				bson.E{Key: "email", Value: 1},
				bson.E{Key: "password", Value: 1},
			}, true),
		),
	}
}
