package models

type Role struct {
	BaseModel `json:"inline"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
}
