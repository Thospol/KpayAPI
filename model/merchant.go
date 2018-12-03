package model

import "github.com/globalsign/mgo/bson"

type Merchant struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Username    string        `bson:"username" json:"username"`
	Password    string        `bson:"password" json:"password"`
	BankAccount []BankAccout  `bson:"bank_account" json:"bank_account"`
}
