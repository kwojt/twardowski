package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	tc "github.com/kwojt/twardowski/commons"
	"github.com/kwojt/twardowski/controllers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rand", controllers.GetRandomPoem).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fmt.Println("Port: ", port)

	err := http.ListenAndServe(":"+port, router)
	tc.Check(err)
}
