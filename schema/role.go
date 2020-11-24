package schema

type Role struct {
	ID          string `json:"role_id" bson:"role_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type RoleBodyParam struct {
	Name        string `json:"name" validate:"required" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type DeleteBodyParam struct {
	ID string `json:"role_id" validate:"required" bson:"role_id"`
}
