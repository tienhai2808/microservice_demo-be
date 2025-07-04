package main

import (
	"context"

	"github.com/tienhai2808/microservice_demo-be/common/api"
)

type OrderService interface {
	CreateOrder(context.Context) error

	ValidateOrder(ctx context.Context, payload *api.CreateOrderRequest) error
}

type OrderStore interface {
	Create(context.Context) error
}