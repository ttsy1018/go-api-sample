package mw

import (
	"fmt"
	"net/http"
	"os"

	"github.com/form3tech-oss/jwt-go"
)

// JWT認証
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ヘッダの認証トークン存在確認
		if r.Header["Token"] != nil {
			// トークンの認証
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				secretKey := os.Getenv("JWT_TOKEN_SIGN")
				return []byte(secretKey), nil
			})

			// 認証失敗時
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			// 認証成功時
			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
