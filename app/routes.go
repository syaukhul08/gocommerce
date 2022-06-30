package app

import (
	"github.com/gorilla/mux"
	"github.com/syaukhul08/goCommerce/app/controllers"
)

func (server *Server) InitializeRoutes() {

	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("get")
}
