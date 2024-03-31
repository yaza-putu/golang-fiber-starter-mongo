package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt primitive.DateTime `json:"-" form:"-" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"-" form:"-" bson:"updated_at"`
}

type Users []User
