package handlers

import (
	"encoding/json"
	"net/http"
	"order/configs"
	"order/models"
	"order/repository"
	"strconv"
)

type OrderHandler struct {
	Config          *configs.Config
	OrderRepository *repository.OrderRepository
}

type OrderHandlerDeps struct {
	OrderRepository *repository.OrderRepository
}

func NewOrderHandler(router *http.ServeMux, deps *OrderHandlerDeps) *OrderHandler {
	handler := &OrderHandler{
		OrderRepository: deps.OrderRepository,
	}
	router.Handle("POST /order", handler.Create())
	router.Handle("GET /orders", handler.GetAll())
	router.Handle("GET /orders/{id}", handler.GetById())
	return handler
}

func (handler *OrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order models.Order
		json.NewDecoder(r.Body).Decode(&order)

		createdOrder, err := handler.OrderRepository.Create(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createdOrder)
	}
}

func (handler *OrderHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ordersAll, err := handler.OrderRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ordersAll)
	}
}

func (handler *OrderHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		byId, err := handler.OrderRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(byId)
	}
}
