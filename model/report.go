package model

import "github.com/globalsign/mgo/bson"

type Report struct {
	ID             bson.ObjectId          `bson:"_id" json:"id"`
	Date           string                 `bson:"date" json:"date"`
	ProductSelling []ProductSellingReport `bson:"products" json:"products"`
	Accumulate     float64                `bson:"accumulate" json:"accumulate"`
}

type ProductSellingReport struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	Name          string        `bson:"name" json:"name"`
	SellingVolume int           `bson:"selling_volume" json:"selling_volume"`
}

type BuyProduct struct {
	ProductName string `bson:"product_name" json:"product_name"`
	Volume      int    `bson:"volume" json:"volume"`
}
