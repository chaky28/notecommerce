package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chaky28/notecommerce/app/app/dal"
	"github.com/chaky28/notecommerce/app/app/server"
)

func GetHandlers() server.BluePrint {
	bp := server.BluePrint{Prefix: "/api"}

	bp.AddHandler("/products/list", getProducts)

	return bp
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	validRequest := server.ValidRequest{
		Method: "GET",
		Params: []string{"category"},
	}
	if err := server.ValidateRequest(w, r, validRequest); err != nil {
		fmt.Println(err.Error())
		return
	}

	db := dal.GetNotEcommerceDB()

	products, err := db.GetProductsByCatId(r.URL.Query().Get("category"))
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR: getting products"))
		return
	}

	json, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR: transforming products to json"))
	}

	w.Write(json)
}
