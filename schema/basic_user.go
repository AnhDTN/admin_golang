package schema

type BasicUserParam struct {
	Email    string `json:"email" validate:"required" bson:"email"`
	Password string `json:"pass_word" bson:"pass_word"`
}
