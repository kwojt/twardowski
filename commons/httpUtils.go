package commons

import (
	"encoding/json"
	"net/http"
)

const allowControlAllowOrigin = "*"

func Message(success bool, message string) map[string]interface{} {
	return map[string]interface{}{"success": success, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Access-Control-Allow-Origin", allowControlAllowOrigin)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func RespondHtml(w http.ResponseWriter, body string) {
	w.Header().Add("Access-Control-Allow-Origin", allowControlAllowOrigin)
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(body))
}
