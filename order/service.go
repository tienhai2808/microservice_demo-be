package main

import "context"

type orderServiceImpl struct {
	store OrderStore
}

func NewOrderService(store OrderStore) OrderService {
	return &orderServiceImpl{
		store: store,
	}
}

func (s *orderServiceImpl) CreateOrder(context.Context) error {
	return nil
}