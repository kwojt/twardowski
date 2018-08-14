package controller

import (
	"net/http"
	tc "twardowski/commons"
	ts "twardowski/services"
)

func GetRandomPoem(w http.ResponseWriter, r *http.Request) {
	poem := ts.GetRandomPoem()
	resp := tc.Message(true, "success")
	resp["data"] = poem
	tc.Respond(w, resp)
}
