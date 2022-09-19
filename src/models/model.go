package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Beneficiary struct {
	Id            primitive.ObjectID `json:"id" bson:"_id"`
	PSP           string             `json:"psp" bson:"psp"`
	FullName      string             `json:"fullName" bson:"fullName" binding:"required"`
	Age           string             `json:"age" bson:"age" binding:"required"`
	Gender        string             `json:"gender" bson:"gender" binding:"required"`
	Phone         string             `json:"phone" bson:"phone" binding:"required"`
	State         string             `json:"state" bson:"state" binding:"required"`
	MaritalStatus string             `json:"maritalStatus" bson:"maritalStatus" binding:"required"`
	LGA           string             `json:"lga" bson:"lga" binding:"required"`
	Ward          string             `json:"ward" bson:"ward" binding:"required"`
	Picture       []byte             `json:"picture" bson:"picture"`
	PaymentMethod string             `json:"paymentMethod" bson:"paymentMethod" binding:"required"`
	PaymentProof  []byte             `json:"paymentProof" bson:"paymentProof"`
	Mode           string            `bson:"mode"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
}
