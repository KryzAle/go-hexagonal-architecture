package main

import (
	"log"
	"net/http"

	"github.com/charly3pins/eShop/application"
	"github.com/charly3pins/eShop/infrastructure/api"
	"github.com/charly3pins/eShop/infrastructure/postgres"
)

func main() {
	connectionOptions := postgres.ConnectionOptions{
		Host:     "127.0.0.1",
		Port:     "5434",
		User:     "eShop",
		Password: "eShop",
		DbName:   "eshop_db",
	}

	db, err := postgres.NewConnection(connectionOptions)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepository := postgres.UserRepository{Db: db}
	userService := application.UserService{UserRepository: userRepository}
	userHandler := api.UserHandler{UserService: userService}

	orderRepository := postgres.OrderRepository{Db: db}
	orderService := application.OrderService{OrderRepository: orderRepository}
	orderHandler := api.OrderHandler{OrderService: orderService}

	r := api.BuildRouter(userHandler, orderHandler)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
