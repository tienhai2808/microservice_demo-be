package main

import "context"

type orderStoreImpl struct {
	//Thêm DB
}

func NewOrderStore() OrderStore {
	return &orderStoreImpl{}
}

func (r *orderStoreImpl) Create(context.Context) error {
	return nil
}