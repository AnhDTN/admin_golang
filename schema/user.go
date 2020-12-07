package schema

type UserParam struct {
	Id       string `json:"id" bson:"id"`
	Email    string `json:"email" validate:"required" bson:"email"`
	Phone    string `json:"phone" validate:"required" bson:"phone"`
	UserName string `json:"user_name" bson:"user_name"`
}

type LoginBodyParam struct {
	Email    string `json:"email" validate:"required" bson:"email"`
	Password string `json:"pass_word" validate:"required" bson:"pass_word"`
}

type RegisterBodyParam struct {
	Email    string `json:"email" validate:"required,email" bson:"email"`
	Password string `json:"pass_word" validate:"required,min=6,max=30" bson:"pass_word"`
	Phone    string `json:"phone" validate:"required" bson:"phone"`
	UserName string `json:"user_name" validate:"required" bson:"user_name"`
	RoleId   string `json:"role_id" validate:"required" bson:"role_id"`
}

type UserIdParam struct {
	Id string `json:"id" bson:"id"`
}

type RefreshTokenParam struct {
	RefreshToken string `json:"refresh_token,omitempty" bson:"refresh_token"`
}

type UserUpdateParam struct {
	Id           string `json:"id" bson:"id"`
	UpdateTime   string `json:"update_time" bson:"update_time"`
	UserName     string `json:"user_name" validate:"require" bson:"user_name"`
	Email        string `json:"email" validate:"require" bson:"email"`
	Password     string `json:"pass_word" validate:"require" bson:"pass_word"`
	Phone        string `json:"phone" validate:"require" bson:"phone"`
	RoleId       string `json:"role_id" validate:"require" bson:"role_id"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type UserTokenInfo struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type UserUpdateBodyParam struct {
	Id           string `json:"id" bson:"id"`
	Password     string `json:"password,omitempty"`
	RoleID       string `json:"role_id,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
