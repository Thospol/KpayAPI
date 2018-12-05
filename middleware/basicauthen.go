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
		if username == "" || password == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if merchants, err := SDO.FindAll(); err == nil {
			for _, merchantsCollection := range merchants {
				if merchantsCollection.Username == username && merchantsCollection.Password == password {
					return
				}
			}
			c.JSON(http.StatusUnauthorized, map[string]string{"result": "username or password not correct"})
		}
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}
