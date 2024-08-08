package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Username string             `bson:"username" binding:"required" json:"username"`
	Password string             `bson:"password" binding:"required" json:"password"`
	IsAdmin  bool               `bson:"isAdmin" json:"isAdmin"`
}
