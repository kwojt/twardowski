package controllers

import (
	"net/http"

	tc "github.com/kwojt/twardowski/commons"
	ts "github.com/kwojt/twardowski/services"
)

func GetRandomPoem(w http.ResponseWriter, r *http.Request) {
	poem := ts.GetRandomPoem()
	resp := tc.Message(true, "success")
	resp["data"] = poem
	tc.Respond(w, resp)
}
