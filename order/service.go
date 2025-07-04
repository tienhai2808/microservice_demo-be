package main

import (
	"context"
	"log"

	"github.com/tienhai2808/microservice_demo-be/common"
	"github.com/tienhai2808/microservice_demo-be/common/api"
)

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

func (s *orderServiceImpl) ValidateOrder(ctx context.Context, payload *api.CreateOrderRequest) error {
	if len(payload.Items) == 0 {
		return common.ErrNoItems
	}
	
	mergedItems := mergeItemsQuantities(payload.Items)
	log.Println(mergedItems)

	return nil
}

func mergeItemsQuantities(items []*api.ItemWithQuantity) []*api.ItemWithQuantity {
	merged := make([]*api.ItemWithQuantity, 0)

	for _, item := range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.ID == item.ID {
				finalItem.Quantity += item.Quantity
				found = true 
				break
			}
		}
		if !found {
			merged = append(merged, item)
		}
	}

	return merged
}