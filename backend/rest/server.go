package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewMiddleware()

	manager.Use(
		middleware.Cros,
		middleware.Logger,
		middleware.Preflight,
	)

	mux := http.NewServeMux()

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	wrappedMux := manager.WrapMux(mux)
	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server running on", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("error starting the server", err)
		os.Exit(1)
	}
}