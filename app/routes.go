package app

import "github.com/syaukhul08/goCommerce/app/controllers"

func (server *Server) InitializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("get")
}
