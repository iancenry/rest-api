package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/iancenry/go-rest-api/handler"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		cookie, err := r.Cookie("token")
		if err != nil {
            if err == http.ErrNoCookie {
                w.WriteHeader(http.StatusUnauthorized)
                return
            }
            w.WriteHeader(http.StatusBadRequest)
            return
        }


		tokenStr  := cookie.Value
		claims := &handler.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims,  func(token *jwt.Token) (interface{}, error) {
            return handler.JwtKey, nil
        })
		if err != nil || !tkn.Valid {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }


		next.ServeHTTP(w, r)


	})
}

//  curl -X POST http://localhost:8000/login -d '{"username":"admin", "password":"password"}' -H "Content-Type: application/json"
// curl --cookie "token=<your_token>" http://localhost:8000/books