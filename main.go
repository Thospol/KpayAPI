package main

import (
	"fmt"
	"kpay/config"
	"kpay/dataaccessobject"
)

var (
	configs = config.Config{}
	daos    = dataaccessobject.DataAccessObject{}
)

func init() {
	configs.Read()
	daos.Server = configs.Server
	daos.Database = configs.Database
	dbAccess := daos.ConnectDatabase()
	fmt.Println("Connected Database: ", dbAccess)
}

func main() {

}
