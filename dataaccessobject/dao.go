package dataaccessobject

import (
	"log"

	mgo "github.com/globalsign/mgo"
)

type DataAccessObject struct {
	Server   string
	Database string
}

var (
	db *mgo.Database
)

const (
	COLLECTION = "Merchant"
)

func (d *DataAccessObject) ConnectDatabase() *mgo.Database {
	session, err := mgo.Dial(d.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(d.Database)
	return db
}
