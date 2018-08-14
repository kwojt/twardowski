package main

import(
  "fmt"
  "github.com/gorilla/mux"
  "os"
  "net/http"
)

func main() {
  fmt.Println(`pop`)

  router := mux.NewRouter();
  router.HandleFunc("/rand", getRandomPoem).Methods("GET")

  port := os.Getenv("PORT")
  if port == "" {
    port = "80"
  }

  fmt.Println("Port: ", port)

  err := http.ListenAndServe(":" + port, router)
  Check(err)
}