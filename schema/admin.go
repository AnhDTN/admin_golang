package schema

type AdminParam struct {
	Id       string `json:"id" bson:"id"`
	Email    string `json:"email" validate:"require" bson:"email"`
	Phone    string `json:"phone" validate:"require" bson:"phone"`
	UserName string `json:"user_name" bson:"user_name"`
}

type LoginBodyParam struct {
	Email    string `json:"email" validate:"require" bson:"email"`
	Password string `json:"pass_word" validate:"require" bson:"pass_word"`
}

type RegisterBodyParam struct {
	Email    string `json:"email" validate:"require" bson:"email"`
	Password string `json:"pass_word" validate:"require" bson:"pass_word"`
	Phone    string `json:"phone" validate:"require" bson:"phone"`
	UserName string `json:"user_name" bson:"user_name"`
	RoleId   string `json:"role_id" validate:"require" bson:"role_id"`
}

type AdminIdParam struct {
	Id string `json:"id" bson:"id"`
}

type RefreshTokenParam struct {
	RefreshToken string `json:"refresh_token,omitempty" bson:"refresh_token"`
}

type AdminUpdateParam struct {
	Id         string `json:"id" bson:"id"`
	UpdateTime string `json:"update_time" bson:"update_time"`
	UserName   string `json:"user_name" validate:"require" bson:"user_name"`
	Email      string `json:"email" validate:"require" bson:"email"`
	Password   string `json:"pass_word" validate:"require" bson:"pass_word"`
	Phone      string `json:"phone" validate:"require" bson:"phone"`
	RoleId     string `json:"role_id" validate:"require" bson:"role_id"`
}
