package models

type BasicUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"pass_word" bson:"pass_word"`
}
