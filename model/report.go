package model

import "github.com/globalsign/mgo/bson"

type Report struct {
	ID             bson.ObjectId          `bson:"_id" json:"id"`
	IDMerchant     bson.ObjectId          `bson:"id_merchant" json:"id_merchant"`
	Date           string                 `bson:"date" json:"date"`
	ProductSelling []ProductSellingReport `bson:"products" json:"products"`
	Accumulate     float64                `bson:"accumulate" json:"accumulate"`
}

type ProductSellingReport struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	Name          string        `bson:"name" json:"name"`
	SellingVolume int           `bson:"selling_volume" json:"selling_volume"`
}

type AddReport struct {
	ProductName string `bson:"product_name" json:"product_name"`
	Volume      int    `bson:"volume" json:"volume"`
}

type BuyProduct struct {
	IDMerchant  bson.ObjectId `bson:"id_merchant" json:"id_merchant"`
	ProductName string        `bson:"product_name" json:"product_name"`
	Volume      int           `bson:"volume" json:"volume"`
}
