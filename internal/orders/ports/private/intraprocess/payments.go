package intraprocess

import (
	"github.com/ttlzx/monolith-microservice-shop/internal/orders/application"
	"github.com/ttlzx/monolith-microservice-shop/internal/orders/domain/orders"
)

type OrdersInterface struct {
	service application.OrdersService
}

func NewOrdersInterface(service application.OrdersService) OrdersInterface {
	return OrdersInterface{service}
}

func (p OrdersInterface) MarkOrderAsPaid(orderID string) error {
	return p.service.MarkOrderAsPaid(application.MarkOrderAsPaidCommand{orders.ID(orderID)})
}
