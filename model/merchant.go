package model

import "github.com/globalsign/mgo/bson"

type Merchant struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	BankAccount string        `bson:"name" json:"name"`
	Username    string        `bson:"name" json:"name"`
	Password    string        `bson:"name" json:"name"`
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
