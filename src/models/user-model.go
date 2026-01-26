package models

type body struct {
	Height int `json:"height" bson:"height"`
	Weight int `json:"weight" bson:"weight"`
}
type User struct {
	BaseModel `bson:",inline"`
	UID       string `json:"uid" bson:"uid"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	PhotoUrl  string `json:"photo_url" bson:"photo_url"`
	Body      body   `json:"body" bson:"body"`
	// RoleID    string `json:"role_id" bson:"role_id,omitempty"`
}
