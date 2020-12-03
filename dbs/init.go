package dbs

import (
	"Login_Golang/model"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
	"time"
)

var (
	once   sync.Once
	Client *mongo.Client
)

func init() {
	url := "mongodb://localhost:27017/admin_golang"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	once.Do(func() {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
		if err != nil {
			log.Print("Client Connect Error: ", err )
			panic(err)
		}
		Client = client
	},
	)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := Client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Print("Client no connect to server: ",err)
		panic(err)
	}

	database := Client.Database("admin_golang")

	adminCollection := database.Collection("admin")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := adminCollection.InsertOne(ctx, model.User{
		BaseModel: model.BaseModel{
			 Id:        uuid.New().String(),
			TypeId:     "1",
			CreateTime: time.Now().String(),
			UpdateTime: time.Now().String(),
		},
		Name: "Nam Anh",
		Email: "anh.dtn@gmail.com",
		Token: uuid.New().String(),
	})

	if err!= nil {
		log.Print("Insert Admin Error: ",err)
	}

	log.Print("Insert repository: ",res.InsertedID)
}
