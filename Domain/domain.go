package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Title       string             `bson:"title" binding:"required" json:"title"`
	Description string             `bson:"description" binding:"required" json:"description"`
	DueDate     time.Time          `bson:"due_date" binding:"required" json:"due_date"`
	Status      string             `bson:"status" binding:"required" json:"status"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Username string             `bson:"username" binding:"required" json:"username"`
	Password string             `bson:"password" binding:"required" json:"password"`
	IsAdmin  bool               `bson:"isAdmin" json:"isAdmin"`
}
