package impl

import (
	"admin_golang/dbs"
	"admin_golang/models"
	"admin_golang/schema"
	"context"
	"log"
	"time"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"

	"admin_golang/pkg"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/mongo"
)

const adminCollection = "admin"

type AdminRepo struct {
	collection *mongo.Collection
}

func NewAdminRepository() *AdminRepo {
	return &AdminRepo{collection: dbs.Database.Collection(adminCollection)}
}

func (a *AdminRepo) CreateUser(body *schema.RegisterBodyParam) (*models.User, error) {
	var admin models.User
	copier.Copy(&admin, &body)
	admin.Id = uuid.New().String()
	admin.PassWord = pkg.HashAndSalt([]byte(body.Password))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := a.collection.InsertOne(ctx, admin)
	if err != nil {
		log.Print("Create admin error: ", err)
		return nil, err
	}
	return &admin, nil
}

func (a *AdminRepo) DeleteUser(admin *schema.AdminIdParam) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := a.collection.DeleteOne(ctx, admin.Id)
	if err != nil {
		log.Print("Delete User error: ", err)
		return err
	}
	return nil
}

func (a *AdminRepo) UpdateUser(admin *schema.AdminUpdateParam) (*models.User, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result models.User
	copier.Copy(&result, &admin)
	filter := admin.Id
	update := admin
	_, err := a.collection.UpdateOne(cxt, filter, update)
	if err != nil {
		log.Print("Update User error: ", err)
		return nil, err
	}
	return &result, nil
}

func (a *AdminRepo) GetUserById(param *schema.AdminIdParam) (*models.User, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result models.User
	err := a.collection.FindOne(cxt, param.Id).Decode(&result)
	if err != nil {
		log.Print("Get User By Id Error: ", err)
		return nil, err
	}
	return &result, nil
}

func (a *AdminRepo) GetUserByToken(token string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result models.User
	filter := bson.M{"refresh_token": token}
	err := a.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Print("Get User By Token Error: ", err)
		return nil, err
	}
	return &result, nil
}

func (a *AdminRepo) DeleteUserToken(param *schema.AdminIdParam) error {
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := param.Id
	update := bson.D{{"refresh_token", ""}}
	_, err := a.collection.UpdateOne(cxt, filter, update)
	if err != nil {
		log.Print("Delete User By RefreshToken Error: ", err)
		return err
	}
	return nil
}

func (a *AdminRepo) GetAllUser() ([]*models.User, error) {
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var results []*models.User
	cur, err := a.collection.Find(cxt, bson.D{{}})
	if err != nil {
		log.Print("Get All AdminError: ", err)
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var element models.User
		err := cur.Decode(&element)
		if err != nil {
			log.Print("Decode Element Error: ", err)
			return nil, err
		}
		results = append(results, &element)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(context.TODO())
	return results, nil
}

func (a *AdminRepo) RefreshToken(token string) (*models.User, error) {
	panic("implement me")
}
