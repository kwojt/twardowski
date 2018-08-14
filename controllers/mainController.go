package controllers

import (
	"net/http"

	tc "github.com/kwojt/twardowski/commons"
)

const apiMessage = `
Hello! This app uses following API:<br>
===================================<br>
/rand - GETs JSON with random poem<br>`

func DisplayAPI(w http.ResponseWriter, r *http.Request) {
	tc.RespondHtml(w, apiMessage)
}
