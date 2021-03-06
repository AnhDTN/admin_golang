package models

type Admin struct {
	BaseModel `json:"inline"`
	UserName string `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	PassWord string `json:"pass_word,omitempty" bson:"pass_word,omitempty"`
	Phone string `json:"phone,omitempty" bson:"phone,omitempty"`
	RoleId string `json:"role_id,omitempty" bson:"role_id,omitempty"`
	Token string `json:"token,omitempty" bson:"token,omitempty"`
}
