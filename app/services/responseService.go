package services

import (
	"net/http"
	"encoding/json"
)

// APIレスポンス（bodyあり）
func SendResponse(w http.ResponseWriter, response []byte, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// エラーレスポンス作成 (エラーメッセージはstring)
func CreateErrorStringResponse(errMessage string) []byte {
	response := map[string]interface{}{
		"error": errMessage,
	}
	responseBody, _ := json.Marshal(response)

	return responseBody
}