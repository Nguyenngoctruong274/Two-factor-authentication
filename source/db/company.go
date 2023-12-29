package db

import (
	"context"

	"github.com/ioVN/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobDecriptionDB struct {
	ac *mongo.Collection
}

func (a *JobDecriptionDB) Job(ctx context.Context, id primitive.ObjectID) (*JobDecription, error) {

	c := &JobDecription{}

	res := a.ac.FindOne(ctx, bson.M{
		"_id": id,
	}, nil)

	if err := res.Decode(c); err != nil {
		return nil, err
	}

	return c, nil
}

func NewJobDecriptionDB(db *mongo.Database) *JobDecriptionDB {
	return &JobDecriptionDB{
		ac: database.MongoInit(
			db,
			"jobDescription",
			database.MakeMongoIndex(bson.D{
				bson.E{Key: "_id", Value: 1},
				bson.E{Key: "company", Value: 1},
				bson.E{Key: "position", Value: 1},
				bson.E{Key: "age", Value: 1},
			}, true),
		),
	}
}
