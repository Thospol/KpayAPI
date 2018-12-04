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

func (u *DataAccessObject) AddProduct(addproduct *model.AddProduct, merchant model.Merchant) (*model.Merchant, error) {

	if addproduct.NameProduct == "" {
		return nil, errors.New("please require  NameProduct")
	}

	var productMerchant model.Product
	var amountHistory model.AmountHistory

	resultAction := "ราคาเริ่มต้น"
	amountHistory.ID = bson.NewObjectId()
	amountHistory.Amount = addproduct.Amount
	amountHistory.Action = resultAction

	productMerchant.ID = bson.NewObjectId()
	productMerchant.NameProduct = addproduct.NameProduct
	productMerchant.Amount = addproduct.Amount
	productMerchant.AmountChange = append(productMerchant.AmountChange, amountHistory)
	merchant.Products = append(merchant.Products, productMerchant)

	err := db.C(COLLECTION).UpdateId(merchant.ID, &merchant)
	fmt.Printf("%#v\n", merchant)

	return &merchant, err
}

func (u *DataAccessObject) DeleteProductMerchant(product_id string, merchant model.Merchant) (model.Merchant, error) {

	var product []model.Product
	for _, merchants := range merchant.Products {
		if merchants.ID == bson.ObjectIdHex(product_id) {
			fmt.Println("Delete BankAccount.ID =", merchants.ID)
		} else {
			product = append(product, merchants)
		}
	}
	merchant.Products = product
	err := db.C(COLLECTION).UpdateId(merchant.ID, &merchant)
	fmt.Printf("%#v\n", merchant)

	return merchant, err
}

func (u *DataAccessObject) UpdateProductMerchant(product_id string, updateProductMerchant *model.UpdateProduct, merchant model.Merchant) (*model.Merchant, error) {
	if updateProductMerchant.Amount == 0 {
		return nil, errors.New("please require  Amount")
	}
	var productMerchant []model.Product
	var amountHistory model.AmountHistory
	resultAction := "ราคาเปลี่ยนแปลงแล้ว"
	amountHistory.ID = bson.NewObjectId()
	amountHistory.Action = resultAction
	amountHistory.Amount = updateProductMerchant.Amount

	for _, productMerchants := range merchant.Products {
		if productMerchants.ID == bson.ObjectIdHex(product_id) {
			productMerchants.Amount = updateProductMerchant.Amount
			productMerchants.AmountChange = append(productMerchants.AmountChange, amountHistory)

			productMerchant = append(productMerchant, productMerchants)
		} else {
			productMerchant = append(productMerchant, productMerchants)
		}
	}
	merchant.Products = productMerchant

	err := db.C(COLLECTION).UpdateId(merchant.ID, &merchant)
	fmt.Printf("%#v\n", merchant)

	return &merchant, err
}
