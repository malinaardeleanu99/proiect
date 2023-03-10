package middleware

import (
	"example/API/initializers"
	"example/API/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	//get the cookie off req
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//atach to req
		c.Set("user", user)

		//continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//check exp

	//find the user with token sub

	//atach to req

	//continue
	c.Next()
}
