package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system.
// @Description User model
// @ID user
// @Property id string true "User ID"
// @Property name string true "User Name"
// @Property email string true "User Email"
// @Property age int true "User Age"
// @Property isDeleted bool true "Is User Deleted"
// @Property createdAt string true "User Creation Time"
// @Property updatedAt string true "User Last Updated Time"
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Age       int                `json:"age" bson:"age"`
	IsDeleted bool               `json:"isDeleted" bson:"isDeleted"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// ErrorResponse represents the structure of the error response returned by the API.
// @Description Error response model
// @ID error-response
// @Property message string true "Error Message"
// @Property code int false "Error Code"
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
