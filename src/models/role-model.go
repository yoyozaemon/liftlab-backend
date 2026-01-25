package models

type Role struct {
	BaseModel
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
