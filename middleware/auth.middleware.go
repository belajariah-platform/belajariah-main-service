package middleware

import (
	"belajariah-main-service/utils"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const pubKeyPath = "keys/app.rsa.pub"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var verifyKey *rsa.PublicKey
		var token *jwt.Token

		verifyBytes, err := ioutil.ReadFile(pubKeyPath)
		if err == nil {
			verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
			jwtToken := c.Request.Header.Get("Token")

			// validate the token
			token, err = jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				// since we only use the one private key to sign the tokens,
				// we also only use its public counter part to verify
				return verifyKey, nil
			})

			// branch out into the possible error from signing
			switch err.(type) {
			case nil: // no error
				if !token.Valid { // but may still be invalid
					err = fmt.Errorf("Invalid token")
				}

			case *jwt.ValidationError: // something was wrong during the validation
				vErr := err.(*jwt.ValidationError)

				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					err = fmt.Errorf("Token Expired")

				default:
					err = fmt.Errorf("Token ValidationError error: %v", vErr.Errors)
				}

			default: // something else went wrong
				err = fmt.Errorf("Token parse error: %v", err)
			}
		}

		if err == nil {
			c.Next()
		} else {
			utils.PushLogf(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
	}
}
