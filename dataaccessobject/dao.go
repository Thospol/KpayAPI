package dataaccessobject

import (
	"errors"
	"fmt"
	"kpay/helper"
	"kpay/model"
	"log"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type DataAccessObject struct {
	Server   string
	Database string
}

var (
	db *mgo.Database
)

const (
	COLLECTION = "merchants"
)

func (d *DataAccessObject) ConnectDatabase() *mgo.Database {
	session, err := mgo.Dial(d.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(d.Database)
	return db
}

func (u *DataAccessObject) Insert(merchant model.Merchant) error {
	err := db.C(COLLECTION).Insert(merchant)
	fmt.Printf("%#v\n", merchant)
	return err
}

func (u *DataAccessObject) Register(register *model.Register) (*model.Merchant, error) {
	if register.Name == "" {
		return nil, errors.New("please require  merchantName")
	}
	if register.BankAccount == "" {
		return nil, errors.New("please require bankAccountofmerchant")
	}

	var merchant model.Merchant
	var bankAccountMerchant model.BankAccout

	bankAccountMerchant.ID = bson.NewObjectId()
	bankAccountMerchant.AccountNumber = register.BankAccount
	bankAccountMerchant.Balance = register.Balance

	merchant.ID = bson.NewObjectId()
	merchant.Name = register.Name
	merchant.Username = helper.RandomUsername()
	merchant.Password = helper.RandomPassword()

	merchant.BankAccount = append(merchant.BankAccount, bankAccountMerchant)

	err := db.C(COLLECTION).Insert(merchant)
	fmt.Printf("%#v\n", merchant)
	return &merchant, err
}

func (u *DataAccessObject) FindAll() ([]model.Merchant, error) {
	var merchants []model.Merchant
	err := db.C(COLLECTION).Find(bson.M{}).All(&merchants)
	fmt.Printf("%#v\n", merchants)
	return merchants, err
}

func (u *DataAccessObject) FindById(id string) (model.Merchant, error) {
	var merchant model.Merchant
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&merchant)
	fmt.Printf("%#v\n", merchant)
	return merchant, err
}

func (u *DataAccessObject) Update(merchant model.Merchant) error {
	err := db.C(COLLECTION).UpdateId(merchant.ID, &merchant)
	fmt.Printf("%#v\n", merchant)
	return err
}
