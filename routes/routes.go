package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	TodoRoutes(r)
	UserRoutes(r)
	ProductRoutes(r)
	AuthRoutes(r)
	TransactionRoutes(r)
	CartRoutes(r)
	// Call UserRoutes function here ...
}
