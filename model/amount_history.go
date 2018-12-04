package model

import "github.com/globalsign/mgo/bson"

type AmountHistory struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Action string        `bson:"action" json:"action"`
	Amount float64       `bson:"amount" json:"amount"`
}
