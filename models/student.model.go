package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name             string             `json:"name" bson:"name"`
	FatherFullName   string             `json:"father_full_name" bson:"father_full_name"`
	FatherFirstName  string             `json:"father_first_name" bson:"father_first_name"`
	FatherMiddleName string             `json:"father_middle_name" bson:"father_middle_name"`
	FatherLastName   string             `json:"father_last_name" bson:"father_last_name"`
	DOB              time.Time          `json:"dob" bson:"dob"`
	Age              int                `json:"age" bson:"age"`
	Gender           string             `json:"gender" bson:"gender"`
}
