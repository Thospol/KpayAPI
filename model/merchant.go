package model

import "github.com/globalsign/mgo/bson"

type Merchant struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Username    string        `bson:"name" json:"username"`
	Password    string        `bson:"name" json:"password"`
	BankAccount []BankAccout  `bson:"name" json:"bank_account"`
}

type BankAccout struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	AccountNumber string        `bson:"account_number" json:"account_number"`
	Balance       float64       `bson:"balance" json:"balance"`
}

type Product struct {
	ID           bson.ObjectId   `bson:"_id" json:"id"`
	NameProduct  string          `bson:"name_product" json:"name_product"`
	Amount       float64         `bson:"amount" json:"amount"`
	AmountChange []AmountHistory `bson:"amount_change" json:"amount_change"`
}

type AmountHistory struct {
	Action string  `bson:"action" json:"action"`
	Amount float64 `bson:"amount" json:"amount"`
}
