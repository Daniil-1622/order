package main

import (
	"fmt"
	"net/http"
	"order/configs"
	db2 "order/db"
	"order/handlers"
	"order/repository"
)

func main() {
	conf := configs.LoadConfig()
	db := db2.NewDb(conf)
	router := http.NewServeMux()

	// Repository
	orderRepository := repository.NewOrderRepository(db)
	// Handler
	handlers.NewOrderHandler(router, &handlers.OrderHandlerDeps{
		OrderRepository: orderRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
