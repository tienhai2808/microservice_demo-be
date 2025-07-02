package main

import (
	"net/http"
	"github.com/tienhai2808/microservice_demo-be/common/api"
)

type httpHandler struct {
	client api.OrderServiceClient
}

func NewHttpHandler(client api.OrderServiceClient) *httpHandler {
	return &httpHandler{
		client: client,
	}
}

func (h *httpHandler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *httpHandler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	h.client.CreateOrder(r.Context(), &api.CreateOrderRequest{
		CustomerID: customerID,
	})
}
