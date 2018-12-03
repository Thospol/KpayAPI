package model

import "github.com/globalsign/mgo/bson"

type BankAccout struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	AccountNumber string        `bson:"account_number" json:"account_number"`
	Balance       float64       `bson:"balance" json:"balance"`
}
