package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chaky28/notecommerce/app/app/dal"
	"github.com/chaky28/notecommerce/app/app/server"
)

func GetHandlers() server.BluePrint {
	bp := server.BluePrint{Prefix: "/api"}

	// ---------- GET handle functions ----------
	bp.AddHandler("/products/list", getProducts)
	bp.AddHandler("/instalments/list", getInstalments)
	bp.AddHandler("/currencies/list", getCurrencies)

	// ---------- PUT handle functions ----------
	bp.AddHandler("/products/insert", insertProducts)
	bp.AddHandler("/instalments/insert", insertInstalment)
	bp.AddHandler("currencies/insert", insertCurrency)

	// ---------- POST handle functions ----------

	// ---------- DELETE handle functions ----------

	return bp
}

// ------------------------------- GET endpoints -------------------------------

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

func getInstalments(w http.ResponseWriter, r *http.Request) {
	validRequest := server.ValidRequest{
		Method: "GET",
		Params: []string{"instalment"},
	}
	if err := server.ValidateRequest(w, r, validRequest); err != nil {
		fmt.Println(err.Error())
		return
	}

	db := dal.GetNotEcommerceDB()

	instalment, err := db.GetInstalmentById(r.URL.Query().Get("instalment"))
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR: getting instalment"))
		return
	}

	json, err := json.Marshal(instalment)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR: transforming instalment to json"))
	}

	w.Write(json)
}

func getCurrencies(w http.ResponseWriter, r *http.Request) {
	validRequest := server.ValidRequest{
		Method: "GET",
		Params: []string{"currency"},
	}
	if err := server.ValidateRequest(w, r, validRequest); err != nil {
		fmt.Println(err.Error())
		return
	}

	db := dal.GetNotEcommerceDB()

	currency, err := db.GetCurrencyById(r.URL.Query().Get("currency"))
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR: getting currency"))
		return
	}

	json, err := json.Marshal(currency)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR: transforming currency to json"))
	}

	w.Write(json)
}

// ------------------------------- PUT endpoints -------------------------------

func insertProducts(w http.ResponseWriter, r *http.Request) {
	validRequest := server.ValidRequest{
		Method: "PUT",
		Headers: []server.Header{
			{Name: "content-type", Value: "application/json"},
		},
	}
	if err := server.ValidateRequest(w, r, validRequest); err != nil {
		fmt.Println(err)
		return
	}

	var js []dal.Product

	err := json.NewDecoder(r.Body).Decode(&js)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ERROR: converting json. Bad format?"))
		return
	}

	conn := dal.GetNotEcommerceDB()
	for i, prod := range js {
		err := conn.InsertProduct(prod)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR: inserting product N°" + strconv.Itoa(i)))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func insertInstalment(w http.ResponseWriter, r *http.Request) {
	validRequest := server.ValidRequest{
		Method: "PUT",
		Headers: []server.Header{
			{Name: "content-type", Value: "application/json"},
		},
	}
	if err := server.ValidateRequest(w, r, validRequest); err != nil {
		fmt.Println(err)
		return
	}

	var js dal.Instalment

	err := json.NewDecoder(r.Body).Decode(&js)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ERROR: converting json. Bad format?"))
		return
	}

	conn := dal.GetNotEcommerceDB()
	err = conn.InsertInstalment(js)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR: inserting new instalment"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func insertCurrency(w http.ResponseWriter, r *http.Request) {
	validRequest := server.ValidRequest{
		Method: "PUT",
		Headers: []server.Header{
			{Name: "content-type", Value: "application/json"},
		},
	}
	if err := server.ValidateRequest(w, r, validRequest); err != nil {
		fmt.Println(err)
		return
	}

	var js dal.Currency

	err := json.NewDecoder(r.Body).Decode(&js)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ERROR: converting json. Bad format?"))
		return
	}

	conn := dal.GetNotEcommerceDB()
	err = conn.InsertCurrency(js)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR: inserting new currency"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// ------------------------------- POST endpoints -------------------------------
