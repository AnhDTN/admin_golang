package impl

import (
	"admin_golang/dbs"
	"admin_golang/models"
	"admin_golang/repo"
	"admin_golang/schema"
	"context"
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"

	"go.mongodb.org/mongo-driver/mongo"
)

const roleCollection = "role"

type RoleRepo struct {
	collection *mongo.Collection
}

func NewRoleRepository() repo.RoleRepository {
	return &RoleRepo{collection: dbs.Database.Collection(roleCollection)}
}

func (r *RoleRepo) CreateRole(role *schema.RoleBodyParam) (*models.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var newRole models.Role
	copier.Copy(&newRole, role)
	newRole.Id = uuid.New().String()
	_, err := r.collection.InsertOne(ctx, newRole)
	if err != nil {
		log.Print("Create Role Error: ", err)
		return nil, err
	}
	return &newRole, nil
}

func (r *RoleRepo) DeleteRole(role *schema.DeleteBodyParam) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"role_id": role.ID}
	_, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Print("Delete Role Error:", err)
		return err
	}
	return nil
}

func (r *RoleRepo) UpdateRole(role *schema.Role) error {
	var payload map[string]interface{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	data, err := json.Marshal(role)
	if err != nil {
		log.Print("Marshal Fail: ", err)
		return err
	}
	json.Unmarshal(data, &payload)
	filter := bson.D{{"role_id", role.ID}}
	change := bson.D{{"$set", payload}}
	_, err = r.collection.UpdateOne(ctx, filter, change)
	if err != nil {
		log.Print("Update Role Error: ", err)
		return err
	}
	return nil
}

func (r *RoleRepo) GetRoleById(id string) (*models.Role, error) {
	var result models.Role
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := bson.M{"role_id": id}
	err := r.collection.FindOne(ctx, query).Decode(&result)
	if err != nil {
		log.Print("Get Role By Id Error: ", err)
		return nil, err
	}
	return &result, nil
}
