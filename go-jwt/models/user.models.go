package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	FirstName    string             `bson:"first_name" json:"first_name,omitempty" validate:"required,min=2,max=100"`
	LastName     string             `bson:"last_name" json:"last_name,omitempty" validate:"required,min=2,max=100"`
	Password     string             `bson:"password" json:"password,omitempty" validate:"required,min=8"`
	Email        string             `bson:"email" json:"email,omitempty" validate:"email,required"`
	Phone        string             `bson:"phone" json:"phone,omitempty" validate:"required,min=10"`
	Token        string             `bson:"token" json:"token,omitempty"`
	UserType     string             `bson:"user_type" json:"user_type,omitempty" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken string             `bson:"refresh_token" json:"refresh_token,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
	UserId       string             `bson:"user_id" json:"user_id,omitempty"`
}

