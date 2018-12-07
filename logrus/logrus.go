package logrus

import (
	"bytes"
	"fmt"
	"kpay/model"
	"strings"

	logr "github.com/sirupsen/logrus"
)

type LogFormat struct {
	TimestampFormat string
}

func (f *LogFormat) Format(entry *logr.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteByte('[')
	b.WriteString(strings.ToUpper(entry.Level.String()))
	b.WriteString("]:")
	b.WriteString(entry.Time.Format(f.TimestampFormat))

	if entry.Message != "" {
		b.WriteString(" - ")
		b.WriteString(entry.Message)
	}

	if len(entry.Data) > 0 {
		b.WriteString(" || ")
	}
	for key, value := range entry.Data {
		b.WriteString(key)
		b.WriteByte('=')
		b.WriteByte('{')
		fmt.Fprint(b, value)
		b.WriteString("}, ")
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func LogrusMerchantByID(merchant model.Merchant) {
	logr.WithFields(logr.Fields{
		"\nmerchant.ID":          merchant.ID,
		"\nmerchant.Name":        merchant.Name,
		"\nmerchant.Username":    merchant.Username,
		"\nmerchant.Password":    merchant.Password,
		"\nmerchant.BankAccount": merchant.BankAccount,
		"\nmerchant.Products":    merchant.Products,
		"\nmerchant.Report":      merchant.Report,
	}).Info("Merchant")
}

func LogrusMerchantAll(merchant []model.Merchant) {

	for _, merchantList := range merchant {
		logr.WithFields(logr.Fields{
			"\nmerchant.ID":          merchantList.ID,
			"\nmerchant.Name":        merchantList.Name,
			"\nmerchant.Username":    merchantList.Username,
			"\nmerchant.Password":    merchantList.Password,
			"\nmerchant.BankAccount": merchantList.BankAccount,
			"\nmerchant.Products":    merchantList.Products,
			"\nmerchant.Report":      merchantList.Report,
		}).Info("Merchant")
	}
}

func LogrusProductOfMerchantFindList(product []model.Product) {

	for _, prodcutList := range product {
		logr.WithFields(logr.Fields{
			"\nprodcutList.IDMerchant":   prodcutList.IDMerchant,
			"\nprodcutList.ID":           prodcutList.ID,
			"\nprodcutList.NameProduct":  prodcutList.NameProduct,
			"\nprodcutList.Amount":       prodcutList.Amount,
			"\nprodcutList.Volume":       prodcutList.Volume,
			"\nprodcutList.AmountChange": prodcutList.AmountChange,
		}).Info("Merchant -> Prodcut")
	}
}

func LogrusfMerchantPOST(merchant *model.Merchant) {

	logr.WithFields(logr.Fields{
		"\nmerchant.ID":          merchant.ID,
		"\nmerchant.Name":        merchant.Name,
		"\nmerchant.Username":    merchant.Username,
		"\nmerchant.Password":    merchant.Password,
		"\nmerchant.BankAccount": merchant.BankAccount,
	}).Info("Merchant")
}

func LogrusProduct(merchant model.Merchant) {

	logr.WithFields(logr.Fields{
		"\nmerchant.Products": merchant.Products,
	}).Info("Merchant -> Prodcut")
}

func LogrusProductPOSTPUT(merchant *model.Merchant) {

	logr.WithFields(logr.Fields{
		"\nmerchant.Products": merchant.Products,
	}).Info("Merchant -> Prodcut")
}

func LogrusBuyMerchant(merchant *model.Merchant) {
	logr.WithFields(logr.Fields{
		"\nmerchant.ID":          merchant.ID,
		"\nmerchant.Name":        merchant.Name,
		"\nmerchant.Username":    merchant.Username,
		"\nmerchant.Password":    merchant.Password,
		"\nmerchant.BankAccount": merchant.BankAccount,
		"\nmerchant.Products":    merchant.Products,
		"\nmerchant.Report":      merchant.Report,
	}).Info("Merchant")
}

func LogrusReportAll(report []model.Report) {
	for _, reportList := range report {
		logr.WithFields(logr.Fields{
			"\nreport.Date":           reportList.Date,
			"\nreport.ID":             reportList.ID,
			"\nreport.ProductSelling": reportList.ProductSelling,
			"\nreport.Accumulate":     reportList.Accumulate,
		}).Info("Merchant -> Report")
	}
}

func LogrusReportPOST(merchant *model.Merchant) {
	for _, listMerchantReport := range merchant.Report {
		logr.WithFields(logr.Fields{
			"\nreport.Date":           listMerchantReport.Date,
			"\nreport.ID":             listMerchantReport.ID,
			"\nreport.ProductSelling": listMerchantReport.ProductSelling,
			"\nreport.Accumulate":     listMerchantReport.Accumulate,
		}).Info("Merchant -> Report")
	}
}

func LogrusUserPOST(user *model.User) {

	logr.WithFields(logr.Fields{
		"\nuser.ID":        user.ID,
		"\nuser.FirstName": user.FirstName,
		"\nuser.LastName":  user.LastName,
		"\nuser.Username":  user.Username,
		"\nuser.Password":  user.Password,
		"\nuser.IDcard":    user.IDcard,
		"\nuser.Age":       user.Age,
		"\nuser.Email":     user.Email,
		"\nuser.Tel":       user.Tel,
	}).Info("Merchant -> User")
}
