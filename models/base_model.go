package models

import (
	"github.com/google/uuid"
	"time"
)

type BaseModel struct {
	Id         string `json:"id,omitempty" bson:"id,omitempty"`
	CreateTime string `json:"create_time,omitempty" bson:"create_time,omitempty"`
	UpdateTime string `json:"update_time,omitempty" bson:"update_time,omitempty"`
}

func (model *BaseModel) BaseModelCreate() error {
	model.Id = uuid.New().String()
	model.CreateTime = time.Now().String()
	model.UpdateTime = time.Now().String()
	return nil
}
