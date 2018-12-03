package model

type AmountHistory struct {
	Action string  `bson:"action" json:"action"`
	Amount float64 `bson:"amount" json:"amount"`
}
