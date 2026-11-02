package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	prdHandler "ecommerce/rest/handlers/product"
	usrHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middleware"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
	//fmt.Printf("%+v", cnf.DB)
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDb(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//repo
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	//domain
	usrSvc := user.NewService(userRepo)
	prdctSvc := product.NewService(productRepo)

	middleware := middleware.NewMiddlewares(cnf)

	//handler
	productHandler := prdHandler.NewHandler(middleware, prdctSvc)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}