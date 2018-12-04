package model

type Register struct {
	Name        string  `bson:"name" json:"name"`
	BankAccount string  `bson:"bank_account" json:"bank_account"`
	Balance     float64 `bson:"balance" json:"balance"`
}
