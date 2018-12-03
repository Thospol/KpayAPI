package model

import "github.com/globalsign/mgo/bson"

type Product struct {
	ID           bson.ObjectId   `bson:"_id" json:"id"`
	NameProduct  string          `bson:"name_product" json:"name_product"`
	Amount       float64         `bson:"amount" json:"amount"`
	AmountChange []AmountHistory `bson:"amount_change" json:"amount_change"`
}
