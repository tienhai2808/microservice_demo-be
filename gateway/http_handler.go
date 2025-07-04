package main

import (
	"errors"
	"net/http"

	"github.com/tienhai2808/microservice_demo-be/common"
	"github.com/tienhai2808/microservice_demo-be/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	var items []*api.ItemWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	order, err := h.client.CreateOrder(r.Context(), &api.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})

	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}

		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return 
	}

	common.WriteJSON(w, http.StatusOK, order)
}

func validateItems(items []*api.ItemWithQuantity) error {
	if len(items) == 0 {
		return common.ErrNoItems
	}

	for _, i := range items {
		if i.ID == "" {
			return errors.New("Item ID is required")
		}

		if i.Quantity <= 0 {
			return errors.New("Item must has a valid quantity")
		}
	}

	return nil
}
