package commands

import (
	"context"

	"github.com/stackus/ftgogo/delivery/internal/domain"
)

type CancelDelivery struct {
	OrderID string
}

type CancelDeliveryHandler struct {
	deliveryRepo domain.DeliveryRepository
	courierRepo  domain.CourierRepository
}

func NewCancelDeliveryHandler(deliveryRepo domain.DeliveryRepository, courierRepo domain.CourierRepository) CancelDeliveryHandler {
	return CancelDeliveryHandler{
		deliveryRepo: deliveryRepo,
		courierRepo:  courierRepo,
	}
}

func (h CancelDeliveryHandler) Handle(ctx context.Context, cmd CancelDelivery) error {
	delivery, err := h.deliveryRepo.Find(ctx, cmd.OrderID)
	if err != nil {
		return err
	}

	if delivery.AssignedCourierID != "" {
		courier, err := h.courierRepo.Find(ctx, delivery.AssignedCourierID)
		if err != nil {
			return err
		}

		courier.CancelDelivery(delivery.DeliveryID)

		err = h.courierRepo.Update(ctx, courier.CourierID, courier)
		if err != nil {
			return err
		}
	}

	delivery.Cancel()

	return h.deliveryRepo.Update(ctx, delivery.DeliveryID, delivery)
}
