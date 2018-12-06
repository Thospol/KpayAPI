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
		merchant, err := SDO.FindById(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{"result": "Not Found!!!"})
		}
		if merchant.Username == username && merchant.Password == password {
			return
		}
		c.JSON(http.StatusUnauthorized, map[string]string{"result": "username or password not correct"})
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}
