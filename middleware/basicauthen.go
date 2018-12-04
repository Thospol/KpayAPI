package middleware

import (
	"kpay/dataaccessobject"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	SDO = dataaccessobject.DataAccessObject{}
)

func BasicAuthenMerchant(c *gin.Context) {

	username, password, ok := c.Request.BasicAuth()
	if ok {
		if merchants, err := SDO.FindAll(); err == nil {
			for _, merchantsCollection := range merchants {
				if merchantsCollection.Username == username && merchantsCollection.Password == password {
					return
				}
			}
		}
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}
