package logics

import (
	"encoding/json"
	"fmt"
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
	Signin(w http.ResponseWriter, r *http.Request) (token string, err error)
	Signup(w http.ResponseWriter, r *http.Request) (token string, err error)
}

type AuthLogic struct {
	ur repositories.UserRepository
}

// ログインパラメータ
type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 会員登録パラメータ
type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// struct to interface
func NewAuthLogicer() AuthLogicer {
	return &AuthLogic{}
}

////////// authインターフェースを満たすauth構造体のメソッド

func (al AuthLogic) Signin(w http.ResponseWriter, r *http.Request) (token string, err error) {
	// リクエストbodyの取り出し
	body, _ := ioutil.ReadAll(r.Body)

	// json decode
	var req SigninRequest
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

func (al *AuthLogic) Signup(w http.ResponseWriter, r *http.Request) (token string, err error) {
	// リクエストbodyの取り出し
	body, _ := ioutil.ReadAll(r.Body)

	// json decode
	var req SignupRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		// todo: send error response

		return "", err
	}

	// バリデーション

	// ユニークチェック
	// Emailからユーザリスト取得
	users, err := al.ur.GetAllByEmail(req.Email)
	if err != nil {
		// todo: send error response

		return "", err
	}
	// ユニークチェック
	if len(users) != 0 {
		// todo: send error response

		return "", nil
	}

	// リクエストされたパスワードをハッシュ化
	hashedPassword := al.hashPassword(req.Password)

	// 構造体にリクエストデータを追加
	var createUser models.User
	createUser.Name = req.Name
	createUser.Email = req.Email
	createUser.Password = string(hashedPassword)
	if err := al.ur.CreateUser(&createUser); err != nil {
		// todo: send error response

		return "", err
	}

	// token発行
	token, err = al.createJwtToken(&createUser)
	if err != nil {
		// todo: send error response

		return "", err
	}

	// rensponse
	return token, nil
}

////////// モジュール内privateメソッド

func (al *AuthLogic) hashPassword(password string) []byte {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashedPassword
}

// jwtトークンを作成
func (al *AuthLogic) createJwtToken(user *models.User) (string, error) {
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

// /*
// jwt認証のミドルウェア
// */
// var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("JWT_TOKEN_SIGN")), nil
// 	},
// 	SigningMethod: jwt.SigningMethodHS256,
// })
