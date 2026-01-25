package models

type User struct {
	BaseModel `bson:",inline"`
	UID       string `json:"uid" bson:"uid"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	RoleID    string `json:"role_id" bson:"role_id,omitempty"`
}
