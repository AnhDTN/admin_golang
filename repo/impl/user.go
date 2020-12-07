package impl

import (
	"admin_golang/dbs"
	"admin_golang/models"
	"admin_golang/pkg/error_custom"
	"admin_golang/repo"
	"admin_golang/schema"
	"context"
	"encoding/json"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

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

func NewAdminRepository() repo.UserRepository {
	return &AdminRepo{collection: dbs.Database.Collection(adminCollection)}
}

func (a *AdminRepo) Login(body *schema.LoginBodyParam) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"email": body.Email}
	err := a.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Print("Login Error: ", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(body.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Print("Compare Password Fail: ", err)
		return nil, err
	}
	user.PassWord = ""
	return &user, err
}

func (a *AdminRepo) CreateUser(body *schema.RegisterBodyParam) (*models.User, error) {
	var user models.User
	copier.Copy(&user, &body)
	user.Id = uuid.New().String()
	user.CreateTime = time.Now().UTC().String()
	user.UpdateTime = ""
	user.PassWord = pkg.HashAndSalt([]byte(body.Password))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"email": body.Email}
	err := a.collection.FindOne(ctx, filter)

	if err == nil || err.Err() != mongo.ErrNoDocuments {
		return nil, error_custom.ErrorExistUser.New()
	}
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, errInsert := a.collection.InsertOne(ctx, user)

	if errInsert != nil {
		log.Print("Create user error_custom: ", errInsert)
		return nil, error_custom.ErrorInsertDocument.New()
	}
	user.PassWord = ""
	return &user, nil
}

func (a *AdminRepo) DeleteUser(admin *schema.UserIdParam) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := a.collection.DeleteOne(ctx, admin.Id)
	if err != nil {
		log.Print("Delete User error_custom: ", err)
		return err
	}
	return nil
}

func (a *AdminRepo) UpdateUser(admin *schema.UserUpdateParam) (*models.User, error) {
	var payload map[string]interface{}
	var result models.User
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	copier.Copy(&result, &admin)
	data, _ := json.Marshal(admin)
	json.Unmarshal(data, &payload)
	filter := bson.M{"id": admin.Id}
	update := bson.D{{"$set", payload}}
	_, err := a.collection.UpdateOne(cxt, filter, update)
	if err != nil {
		log.Print("Update User error_custom: ", err)
		return nil, err
	}
	return &result, nil
}

func (a *AdminRepo) UpdateRefreshToken(param *schema.UserUpdateBodyParam) error {
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var result models.User
	copier.Copy(&result, &param)
	filter := bson.M{"id": param.Id}
	update := bson.D{{"$set", param}}
	_, err := a.collection.UpdateOne(cxt, filter, update)
	if err != nil {
		log.Print("Update User Error: ", err)
		return err
	}
	return nil
}

func (a *AdminRepo) GetUserById(param *schema.UserIdParam) (*models.User, error) {
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

func (a *AdminRepo) DeleteUserToken(param *schema.UserIdParam) error {
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
