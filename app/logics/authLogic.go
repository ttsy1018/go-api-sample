package logics

import (
	"encoding/json"
	"io/ioutil"
	"myapp/models"
	"myapp/repositories"
	"net/http"
	"os"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthLogicer interface {
	Login(w http.ResponseWriter, r *http.Request) (token string, err error)
}

type AuthLogic struct {
	ur repositories.UserRepository
}

type LoginRequest struct {
	Email    string `json:email`
	Password string `json:password`
}

func NewAuthLogicer() AuthLogicer {
	return &AuthLogic{}
}

func (al AuthLogic) Login(w http.ResponseWriter, r *http.Request) (token string, err error) {
	// リクエストbodyの取り出し
	body, _ := ioutil.ReadAll(r.Body)

	// json decode
	var req LoginRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		// todo: send error response

		return "", err
	}

	// Emailからユーザ取得
	user, err := al.ur.FindByEmail(req.Email)
	if err != nil {
		// todo: send error response

		return "", err
	}

	// パスワード確認(一致だとnil,不一致だとエラーが返る)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		msg := "メールアドレスまたはパスワードが違います。"

		// todo: send error resonse

		return msg, err
	}

	// token発行
	token, err = al.createJwtToken(&user)
	if err != nil {
		// todo: send error response

		return "", err
	}

	// rensponse
	return token, nil
}

// jwtトークンを作成
func (al *AuthLogic) createJwtToken(user *models.User) (string, error) {
	// header
	token := jwt.New(jwt.SigningMethodHS256)

	// claims
	token.Claims = jwt.MapClaims{
		"id":  user.ID,
		"aud": user.Email + user.Name,
		"exp": time.Now().Add(time.Hour * 24), // 24時間後
	}

	// sign
	secretKey := os.Getenv("JWT_TOKEN_SIGN")

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
