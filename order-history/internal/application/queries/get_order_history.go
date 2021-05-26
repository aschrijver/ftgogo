package queries

import (
	"context"

	"github.com/stackus/ftgogo/order-history/internal/domain"
)

type GetOrderHistory struct {
	OrderID string
}

type GetOrderHistoryHandler struct {
	repo domain.OrderHistoryRepository
}

func NewGetOrderHistoryHandler(orderHistoryRepo domain.OrderHistoryRepository) GetOrderHistoryHandler {
	return GetOrderHistoryHandler{repo: orderHistoryRepo}
}

func (h GetOrderHistoryHandler) Handle(ctx context.Context, query GetOrderHistory) (*domain.OrderHistory, error) {
	return h.repo.Find(ctx, query.OrderID)
}
