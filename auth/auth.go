package auth

import (
	"log"
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

func makeTimestamp() int64 {
	//add 7 days
	expires := time.Now().AddDate(0, 0, 7)
	return expires.UnixNano() / int64(time.Millisecond)
}

// Глобальный секретный ключ
var mySigningKey = []byte("secret123456")

//GetTokenHandler gfbfgn
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: makeTimestamp(),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// Отдаем токен клиенту
	w.Write([]byte(tokenString))
})

var jwtMiddleware = jwtmiddleware.New(
	jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

//MiddlewareHandler fgfgfgfg
var MiddlewareHandler = func(h http.Handler) http.Handler {
	return jwtMiddleware.Handler(h)
}
