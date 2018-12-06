package model

import "github.com/globalsign/mgo/bson"

type Product struct {
	ID           bson.ObjectId   `bson:"_id" json:"id"`
	IDMerchant   bson.ObjectId   `bson:"id_merchant" json:"id_merchant"`
	NameProduct  string          `bson:"name_product" json:"name_product"`
	Amount       float64         `bson:"amount" json:"amount"`
	Volume       int             `bson:"volume" json:"volume"`
	AmountChange []AmountHistory `bson:"amount_change" json:"amount_change"`
}

type AddProduct struct {
	NameProduct string  `bson:"name_product" json:"name_product"`
	Amount      float64 `bson:"amount" json:"amount"`
	Volume      int     `bson:"volume" json:"volume"`
}

type UpdateProduct struct {
	Amount float64 `bson:"amount" json:"amount"`
}
