package impl

import (
	"admin_golang/dbs"
	"admin_golang/models"
	"admin_golang/repo"
	"admin_golang/schema"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

const basicCollection = "basic"

type BasicUserRepo struct {
	collection *mongo.Collection
}

func NewBasicUserRepo() repo.IBasicUserRepo {
	return &BasicUserRepo{collection: dbs.Database.Collection(basicCollection)}
}

func (b BasicUserRepo) Create(param *schema.BasicUserParam) error {
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := b.collection.UpdateOne(cxt, param, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func (b BasicUserRepo) GetAll() ([]*models.BasicUser, error) {
	var results []*models.BasicUser
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := b.collection.Find(cxt, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var element models.BasicUser
		err := cur.Decode(element)
		if err != nil {
			return nil, err
		}
		results = append(results, &element)
	}
	return results, nil
}
