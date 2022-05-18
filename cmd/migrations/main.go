package main

import (
	"github.com/charly3pins/eShop/domain"
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

	// TODO create models for DB and add the FK there
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.Product{})
}
