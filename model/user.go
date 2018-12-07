package model

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	ID              bson.ObjectId     `bson:"_id" json:"id"`
	FirstName       string            `bson:"first_name" json:"first_name" binding:"required"`
	LastName        string            `bson:"last_name" json:"last_name" binding:"required"`
	Username        string            `bson:"username" json:"username" binding:"required"`
	Password        string            `bson:"password" json:"password" binding:"required"`
	IDcard          string            `bson:"idcard" json:"idcard" binding:"required"`
	Age             int64             `bson:"age" json:"age" binding:"required"`
	Email           string            `bson:"email" json:"email" binding:"required"`
	Tel             string            `bson:"tel" json:"tel" binding:"required"`
	UserBankAccount []UserBankAccount `bson:"user_bank_account" json:"user_bank_account"`
}
type UserBankAccount struct {
	BankName      string `bson:"bank_name" json:"bank_name"`
	AccountNumber string `bson:"account_number" json:"account_number"`
	Balance       int64  `bson:"balance" json:"balance"`
}
