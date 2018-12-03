package model

import "github.com/globalsign/mgo/bson"

type Merchant struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Username    string        `bson:"name" json:"username"`
	Password    string        `bson:"name" json:"password"`
	BankAccount []BankAccout  `bson:"name" json:"bank_account"`
}
