package model

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	FirstName  string        `bson:"first_name" json:"first_name"`
	LastName   string        `bson:"last_name" json:"last_name"`
	Username   string        `bson:"username" json:"username"`
	Password   string        `bson:"password" json:"password"`
	IDcard     string        `bson:"idcard" json:"idcard"`
	Age        int64         `bson:"age" json:"age"`
	Email      string        `bson:"email" json:"email"`
	Tel        string        `bson:"tel" json:"tel"`
	BankAccout []BankAccout  `bson:"bank_account" json:"bank_account"`
}
