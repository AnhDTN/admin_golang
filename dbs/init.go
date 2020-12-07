package dbs

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// to set env use viper golang

var (
	once     sync.Once
	Database *mongo.Database
	Client   *mongo.Client
)

func init() {
	url := "mongodb://localhost:27017/admin_golang"
	once.Do(func() {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
		if err != nil {
			log.Print("Database Connect Error: ", err)
			panic(err)
		}

		database := client.Database("admin_golang")
		fmt.Println("INIT")

		Client = client
		client = client
		Database = database
		log.Print("database init")

	},
	)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := Client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Print("Database no connect to server: ", err)
		panic(err)
	}

	//adminCollection := Database.Collection("admin")
	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	//resInsertOne, err := adminCollection.InsertOne(ctx, models.User{
	//	BaseModel: models.BaseModel{
	//		Id:         uuid.New().String(),
	//		CreateTime: time.Now().String(),
	//		UpdateTime: time.Now().String(),
	//	},
	//	UserName:     "Nam Anh",
	//	Email:        "anh.dtn@gmail.com",
	//	RefreshToken: uuid.New().String(),
	//})
	//
	//if err != nil {
	//	log.Print("Insert User Error: ", err)
	//}

	//log.Print("Insert repository: ", resInsertOne.InsertedID)
	//
	//ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//resInsertMany, err := adminCollection.InsertMany(ctx, []interface{}{
	//	models.User{
	//		BaseModel: models.BaseModel{
	//			Id:         uuid.New().String(),
	//			CreateTime: time.Now().String(),
	//			UpdateTime: time.Now().String(),
	//		},
	//		UserName:     "Nam Anh",
	//		Email:        "anh.dtn@gmail.com",
	//		RefreshToken: uuid.New().String(),
	//	},
	//	models.User{
	//		BaseModel: models.BaseModel{
	//			Id:         uuid.New().String(),
	//			CreateTime: time.Now().String(),
	//			UpdateTime: time.Now().String(),
	//		},
	//		UserName:     "Nam Anh",
	//		Email:        "anh.dtn@gmail.com",
	//		RefreshToken: uuid.New().String(),
	//	},
	//})

	//if err != nil {
	//	log.Print("Insert many Error: ", err)
	//}
	//
	//log.Print("Insert Many Success: ", resInsertMany)

}
