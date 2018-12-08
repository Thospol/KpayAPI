package dataaccessobject

import (
	"errors"
	"fmt"
	"kpay/helper"
	"kpay/model"
	"log"
	"time"

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
	COLLECTION2 = "users2"
	COLLECTION  = "merchants"
)

func (d *DataAccessObject) ConnectDatabase() *mgo.Database {
	session, err := mgo.Dial(d.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(d.Database)
	return db
}

func (u *DataAccessObject) Register(register *model.Register, merchantCurrentAll []model.Merchant) (*model.Merchant, error) {
	if register.Name == "" {
		return nil, errors.New("please require  merchantName")
	}
	if register.BankAccount == "" {
		return nil, errors.New("please require bankAccountofmerchant")
	}

	for _, Listmerchants := range merchantCurrentAll {
		for _, merchants := range Listmerchants.BankAccount {
			if merchants.AccountNumber == register.BankAccount {
				return nil, errors.New("BankAccount duplicate please specify new")
			}
		}
	}

	var merchant model.Merchant
	var bankAccountMerchant model.BankAccout
	GenetateUsername := helper.StringWithMerchantset(register.Name)
	GenetatePassword := helper.StringWithMerchantset(register.Name)

	bankAccountMerchant.ID = bson.NewObjectId()
	bankAccountMerchant.AccountNumber = register.BankAccount
	bankAccountMerchant.Balance = register.Balance

	merchant.ID = bson.NewObjectId()
	merchant.Name = register.Name
	merchant.Username = GenetateUsername
	merchant.Password = GenetatePassword

	merchant.BankAccount = append(merchant.BankAccount, bankAccountMerchant)

	err := db.C(COLLECTION).Insert(merchant)
	return &merchant, err
}

func (u *DataAccessObject) FindAll() ([]model.Merchant, error) {
	var merchants []model.Merchant
	err := db.C(COLLECTION).Find(bson.M{}).All(&merchants)
	return merchants, err
}

func (u *DataAccessObject) FindById(id string) (model.Merchant, error) {
	var merchant model.Merchant
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&merchant)
	return merchant, err
}

func (u *DataAccessObject) FindProductById(merchant *model.Merchant) []model.Product {
	var products []model.Product
	for _, listproductsInMerchant := range merchant.Products {
		products = append(products, listproductsInMerchant)
	}
	return products
}

func (u *DataAccessObject) Update(merchant model.Merchant) error {
	err := db.C(COLLECTION).UpdateId(merchant.ID, &merchant)
	return err
}

func (u *DataAccessObject) AddProduct(addproduct *model.AddProduct, merchant model.Merchant) (*model.Merchant, error) {

	if addproduct.NameProduct == "" {
		return nil, errors.New("please require  NameProduct")
	}
	if addproduct.Volume == 0 {
		return nil, errors.New("please require  Volume")
	}
	if addproduct.Amount == 0 {
		return nil, errors.New("please require  Amount")
	}

	var productMerchant model.Product
	var amountHistory model.AmountHistory
	if len(merchant.Products) < 5 {
		resultAction := "ราคาเริ่มต้น"
		amountHistory.ID = bson.NewObjectId()
		amountHistory.Amount = addproduct.Amount
		amountHistory.Action = resultAction

		productMerchant.ID = bson.NewObjectId()
		productMerchant.IDMerchant = merchant.ID
		productMerchant.NameProduct = addproduct.NameProduct
		productMerchant.Amount = addproduct.Amount
		productMerchant.Volume = addproduct.Volume
		productMerchant.AmountChange = append(productMerchant.AmountChange, amountHistory)
		merchant.Products = append(merchant.Products, productMerchant)
	} else {
		return nil, errors.New("product maximum 5 product To Merchant")
	}

	err := db.C(COLLECTION).UpdateId(merchant.ID, &merchant)

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

	return &merchant, err
}

func (u *DataAccessObject) ResetReport(report *model.AddReport, merchant model.Merchant) (*model.Merchant, error) {
	if report.ProductName == "" {
		return nil, errors.New("please require  ProductName")
	}
	if report.Volume == 0 {
		return nil, errors.New("please require  Volume")
	}
	var reports []model.Report

	merchant.Report = reports
	err := db.C(COLLECTION).UpdateId(merchant.ID, &merchant)

	fmt.Printf("%#v\n", merchant)
	return &merchant, err
}

func (u *DataAccessObject) AddReport(report *model.AddReport, merchant model.Merchant) (*model.Merchant, error) {
	if report.ProductName == "" {
		return nil, errors.New("please require  ProductName")
	}
	if report.Volume == 0 {
		return nil, errors.New("please require  Volume")
	}

	var reports model.Report
	var productListSelling model.ProductSellingReport
	hasProduct := false
	totalBalance := 0.0
	var idMerchant bson.ObjectId
	var products []model.Product
	for _, listmerchant := range merchant.Products {
		if listmerchant.NameProduct == report.ProductName {
			hasProduct = true
			if listmerchant.Volume < report.Volume {
				return nil, errors.New("Sorry Product you need is not enough")
			}
			productListSelling.ID = bson.NewObjectId()
			productListSelling.Name = report.ProductName
			productListSelling.SellingVolume = report.Volume
			totalBalance = listmerchant.Amount * float64(report.Volume)
			idMerchant = listmerchant.IDMerchant
			listmerchant.Volume = (listmerchant.Volume - report.Volume)
			products = append(products, listmerchant)
		} else {
			products = append(products, listmerchant)
		}
	}
	if hasProduct == false {
		return nil, errors.New("don't have Product in Merchant")
	}
	hasMerchantReportToday := false
	hasProductNameToday := false
	var reportListAll []model.Report
	var sellingvolume []model.ProductSellingReport
	for _, merchantListReport := range merchant.Report {
		if merchantListReport.Date == time.Now().Format("02-01-2006") {
			for _, reportsellinglist := range merchantListReport.ProductSelling {
				if reportsellinglist.Name == report.ProductName {
					hasProductNameToday = true
					reportsellinglist.SellingVolume = reportsellinglist.SellingVolume + report.Volume
					sellingvolume = append(sellingvolume, reportsellinglist)
				} else {
					sellingvolume = append(sellingvolume, reportsellinglist)
				}
			}

			hasMerchantReportToday = true

			if hasProductNameToday == false {
				sellingvolume = append(sellingvolume, productListSelling)
			}
			merchantListReport.Accumulate = merchantListReport.Accumulate + totalBalance
			merchantListReport.ProductSelling = sellingvolume
			merchant.BankAccount[0].Balance = merchant.BankAccount[0].Balance + merchantListReport.Accumulate
			reportListAll = append(reportListAll, merchantListReport)
		} else {
			reportListAll = append(reportListAll, merchantListReport)
		}
	}

	if hasMerchantReportToday == true {
		merchant.Report = reportListAll
		merchant.Products = products
	} else {
		reports.ID = bson.NewObjectId()
		reports.IDMerchant = idMerchant
		reports.Date = time.Now().Format("02-01-2006")
		reports.Accumulate = reports.Accumulate + totalBalance
		reports.ProductSelling = append(reports.ProductSelling, productListSelling)
		merchant.BankAccount[0].Balance = merchant.BankAccount[0].Balance + reports.Accumulate
		merchant.Products = products
		reportListAll = append(reportListAll, reports)
		merchant.Report = reportListAll
	}
	err := db.C(COLLECTION).UpdateId(merchant.ID, &merchant)

	fmt.Printf("%#v\n", merchant)
	return &merchant, err
}

func (u *DataAccessObject) FindReportMerchant(merchant model.Merchant) ([]model.Report, error) {

	var ReportResponse []model.Report
	for _, listMerchantOfReport := range merchant.Report {
		ReportResponse = append(ReportResponse, listMerchantOfReport)
	}
	return ReportResponse, nil
}

func (u *DataAccessObject) FindAllReportMerchant() ([]model.Report, error) {

	var ReportResponse []model.Report
	var merchants []model.Merchant
	err := db.C(COLLECTION).Find(bson.M{}).All(&merchants)

	for _, listMerchants := range merchants {
		for _, listMerchant := range listMerchants.Report {
			ReportResponse = append(ReportResponse, listMerchant)
		}
	}
	return ReportResponse, err
}

func (u *DataAccessObject) BuyProductMerchant(inputBuyProduct *model.BuyProduct, user model.User) (*model.Merchant, error) {
	if inputBuyProduct.IDMerchant == "" {
		return nil, errors.New("please require  IDMerchant")
	}
	if inputBuyProduct.ProductName == "" {
		return nil, errors.New("please require  ProductName")
	}
	if inputBuyProduct.Volume == 0 {
		return nil, errors.New("please require  Volume")
	}
	var merchant model.Merchant
	var err error
	if err = db.C(COLLECTION).FindId(inputBuyProduct.IDMerchant).One(&merchant); err != nil {
		return nil, errors.New("please require  Volume")
	}

	var reports model.Report
	var productListSelling model.ProductSellingReport
	hasProduct := false
	totalBalance := 0.0
	var idMerchant bson.ObjectId
	var products []model.Product
	for _, listmerchant := range merchant.Products {
		if listmerchant.NameProduct == inputBuyProduct.ProductName {
			hasProduct = true
			if listmerchant.Volume < inputBuyProduct.Volume {
				return nil, errors.New("Sorry Product you need is not enough")
			}
			productListSelling.ID = bson.NewObjectId()
			productListSelling.Name = inputBuyProduct.ProductName
			productListSelling.SellingVolume = inputBuyProduct.Volume
			totalBalance = listmerchant.Amount * float64(inputBuyProduct.Volume)
			idMerchant = listmerchant.IDMerchant
			listmerchant.Volume = (listmerchant.Volume - inputBuyProduct.Volume)
			products = append(products, listmerchant)
		} else {
			products = append(products, listmerchant)
		}
	}
	if hasProduct == false {
		return nil, errors.New("don't have Product in Merchant")
	}
	hasMerchantReportToday := false
	hasProductNameToday := false
	var reportListAll []model.Report
	var sellingvolume []model.ProductSellingReport
	for _, merchantListReport := range merchant.Report {
		if merchantListReport.Date == time.Now().Format("02-01-2006") {
			for _, reportsellinglist := range merchantListReport.ProductSelling {
				if reportsellinglist.Name == inputBuyProduct.ProductName {
					hasProductNameToday = true
					reportsellinglist.SellingVolume = reportsellinglist.SellingVolume + inputBuyProduct.Volume
					sellingvolume = append(sellingvolume, reportsellinglist)
				} else {
					sellingvolume = append(sellingvolume, reportsellinglist)
				}
			}

			hasMerchantReportToday = true

			if hasProductNameToday == false {
				sellingvolume = append(sellingvolume, productListSelling)
			}
			merchantListReport.Accumulate = merchantListReport.Accumulate + totalBalance
			merchantListReport.ProductSelling = sellingvolume
			merchant.BankAccount[0].Balance = merchant.BankAccount[0].Balance + merchantListReport.Accumulate
			reportListAll = append(reportListAll, merchantListReport)
		} else {
			reportListAll = append(reportListAll, merchantListReport)
		}
	}

	if hasMerchantReportToday == true {
		merchant.Report = reportListAll
		merchant.Products = products
	} else {
		reports.ID = bson.NewObjectId()
		reports.IDMerchant = idMerchant
		reports.Date = time.Now().Format("02-01-2006")
		reports.Accumulate = reports.Accumulate + totalBalance
		reports.ProductSelling = append(reports.ProductSelling, productListSelling)
		merchant.BankAccount[0].Balance = merchant.BankAccount[0].Balance + reports.Accumulate
		merchant.Products = products
		reportListAll = append(reportListAll, reports)
		merchant.Report = reportListAll
	}
	err = db.C(COLLECTION).UpdateId(merchant.ID, &merchant)

	user.UserBankAccount[0].Balance = user.UserBankAccount[0].Balance - int64(totalBalance)
	err = db.C(COLLECTION2).UpdateId(user.ID, &user)

	return &merchant, err
}

func (u *DataAccessObject) FindAllUser() ([]model.User, error) {
	var users []model.User
	err := db.C(COLLECTION2).Find(bson.M{}).All(&users)
	return users, err
}

func (u *DataAccessObject) FindByIDUser(id string) (model.User, error) {
	var user model.User
	err := db.C(COLLECTION2).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (u *DataAccessObject) InsertUser(userRequest *model.User, user []model.User) (*model.User, error) {

	for _, listOfusers := range user {
		if listOfusers.Username == userRequest.Username || listOfusers.IDcard == userRequest.IDcard {
			return nil, errors.New("Username or IDcard duplicate please specify new")
		}
	}
	userRequest.ID = bson.NewObjectId()

	err := db.C(COLLECTION2).Insert(userRequest)
	return userRequest, err
}

func (u *DataAccessObject) InsertBankAccountOfUser(bankaccountReq *model.UserBankAccount, user model.User, users []model.User) error {

	var err error
	for _, userLists := range users {
		for _, userBankaccList := range userLists.UserBankAccount {
			if userBankaccList.AccountNumber == bankaccountReq.AccountNumber {
				err = errors.New("Please try again duplicate AccountNumber")
				return err
			}
		}
	}
	user.UserBankAccount = append(user.UserBankAccount, *bankaccountReq)
	err = db.C(COLLECTION2).UpdateId(user.ID, &user)
	return err
}

func (u *DataAccessObject) UpdateUser(userreq *model.UpdateUser, user model.User) error {
	if userreq.FirstName != "" {
		user.FirstName = userreq.FirstName
	}
	if userreq.LastName != "" {
		user.LastName = userreq.LastName
	}
	if userreq.Username != "" {
		user.Username = userreq.Username
	}
	if userreq.Password != "" {
		user.Password = userreq.Password
	}
	if userreq.IDcard != "" {
		user.IDcard = userreq.IDcard
	}
	if userreq.Age != 0 {
		user.Age = userreq.Age
	}
	if userreq.Email != "" {
		user.Email = userreq.Email
	}
	if userreq.Tel != "" {
		user.Tel = userreq.Tel
	}
	if userreq.BankName != "" {
		user.UserBankAccount[0].BankName = userreq.BankName
	}
	if userreq.AccountNumber != "" {
		user.UserBankAccount[0].AccountNumber = userreq.AccountNumber
	}
	err := db.C(COLLECTION2).UpdateId(user.ID, &user)
	return err
}

func (u *DataAccessObject) DeleteUser(user model.User) error {
	err := db.C(COLLECTION2).Remove(&user)
	return err
}

func (u *DataAccessObject) DeleteUserBankAccount(userBankAccount *model.UserBankAccount, user model.User) error {
	if userBankAccount.AccountNumber == "" {
		return errors.New("please require AccountNumber")
	}
	var userBankAccountNew []model.UserBankAccount
	for _, userBankAccountList := range user.UserBankAccount {
		if userBankAccountList.AccountNumber == userBankAccount.AccountNumber {
		} else {
			userBankAccountNew = append(userBankAccountNew, userBankAccountList)
		}
	}
	user.UserBankAccount = userBankAccountNew
	err := db.C(COLLECTION2).UpdateId(user.ID, &user)
	return err
}
