package middlewares

import (
	"encoding/base64"
	"fmt"
	"github.com/amoghRZP/web-app/dbconfig"
	"github.com/gin-gonic/gin"
	"strings"
)

type Auth struct {
	Name string
	Pass string
}

/*
Setting table name as "auth"
*/
func (Auth) TableName() string {
	return "auth"
}

func LoginMiddleware(c *gin.Context) {

	var identity Auth
	fmt.Println("Login Middleware Invoked")

	//Getting username and password from header
	Authorization := c.Request.Header.Get("Authorization")
	credentials := strings.TrimSpace(Authorization[len("Basic")+1:])
	credentialsDecoded, _ := base64.StdEncoding.DecodeString(credentials)
	userpass := string(credentialsDecoded)
	s := strings.Split(userpass, ":")
	username, password := s[0], s[1]

	if err := dbconfig.DB.Where("name = ?", username).First(&identity).Error; err != nil || password != identity.Pass {
		c.AbortWithStatus(401)
		fmt.Println(err)
	}

	fmt.Println("User Authorized")

	c.Next()
}
