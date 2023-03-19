package services

import (
	"fmt"
	"myapp/models"
	"net/http"
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/pkg/errors"
)

// jwtトークンを作成
func CreateJwtToken(user *models.User) (string, error) {
	// header
	token := jwt.New(jwt.SigningMethodHS256)

	// claims
	token.Claims = jwt.MapClaims{
		"id":  user.ID,
		"aud": user.Email + user.Name,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 24時間後 (unix時間にしないとうまくいかない)
	}

	// sign
	secretKey := os.Getenv("JWT_TOKEN_SIGN")
	fmt.Println([]byte(secretKey))
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// リクエストヘッダからユーザIDを取得
func GetUserIdByRequestToken(r *http.Request) (int, error) {
	// ヘッダからtokenを取得
	clientToken := r.Header.Get("Token")
	if clientToken == "" {
		return 0, errors.New("not token")
	}

	secretKey := os.Getenv("JWT_TOKEN_SIGN")

	// トークンをparse
	token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("トークンをjwtにparseできません。")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	// クレーム（payloadd）を取得
	claims, claimOk := token.Claims.(jwt.MapClaims)
	if !claimOk || !token.Valid {
		return 0, errors.New("id type not match")
	}

	// クレームからユーザIDを取得
	userId, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("id type not match")
	}

	return int(userId), nil
}
