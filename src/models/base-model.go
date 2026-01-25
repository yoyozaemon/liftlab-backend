package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"-" bson:"updated_at,omitempty"`
	DeletedAt time.Time          `json:"-" bson:"deleted_at,omitempty"`
	IsDeleted bool               `json:"-" bson:"is_deleted,omitempty"`
	CreatedBy string             `json:"-" bson:"created_by,omitempty"`
	UpdatedBy string             `json:"-" bson:"updated_by,omitempty"`
	DeletedBy string             `json:"-" bson:"deleted_by,omitempty"`
}
