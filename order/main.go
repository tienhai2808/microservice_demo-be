package main

import "context"

func main() {
	store := NewOrderStore()
	service := NewOrderService(store)

	service.CreateOrder(context.Background())
}