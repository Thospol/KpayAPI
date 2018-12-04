package model

import "github.com/globalsign/mgo/bson"

type Merchant struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Username    string        `bson:"username" json:"username"`
	Password    string        `bson:"password" json:"password"`
	BankAccount []BankAccout  `bson:"bank_account" json:"bank_account"`
	Products    []Product     `bson:"products" json:"products"`
}

type Register struct {
	Name        string  `bson:"name" json:"name"`
	BankAccount string  `bson:"bank_account" json:"bank_account"`
	Balance     float64 `bson:"balance" json:"balance"`
}

type AddProduct struct {
	NameProduct string  `bson:"name_product" json:"name_product"`
	Amount      float64 `bson:"amount" json:"amount"`
}

type UpdateProduct struct {
	Amount float64 `bson:"amount" json:"amount"`
}
