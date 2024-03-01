package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chaky28/notecommerce/app/app/dal"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello everybody, you %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", home)

	ndb := dal.GetNotEcommerceDB()
	fmt.Println(ndb)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
