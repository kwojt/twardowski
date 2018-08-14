package controllers

import (
	"net/http"

	tc "github.com/kwojt/twardowski/commons"
)

const apiMessage = `Hello! This app uses following API:
===================================
/rand - GETs JSON with random poem`

func DisplayAPI(w http.ResponseWriter, r *http.Request) {
	tc.RespondHtml(w, apiMessage)
}
