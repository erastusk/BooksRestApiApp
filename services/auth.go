package services

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
)

const tokenString = "booksapi"

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(tokenString), nil
	})
}

func CreateToken() string {
	tokenStr := jwt.New(jwt.SigningMethodHS256)
	str, err := tokenStr.SignedString([]byte(tokenString))
	fmt.Println(str)
	if err != nil {
		log.Fatal(err)
	}
	return str
}
