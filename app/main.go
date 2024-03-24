package main

import (
	"fmt"
	"net/http"

	"github.com/chaky28/notecommerce/app/app/api"
	"github.com/chaky28/notecommerce/app/app/server"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello everybody, you %s!", r.URL.Path[1:])
}

func main() {
	//Initialize server
	s := server.Server{
		Port: 8080,
	}

	// Register handler blueprints
	bps := []server.BluePrint{
		api.GetHandlers(),
	}

	//Start server
	s.ListenAndServe(bps)
}
