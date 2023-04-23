package intraprocess

import (
	"github.com/ttlzx/monolith-microservice-shop/internal/orders/app"
	"github.com/ttlzx/monolith-microservice-shop/internal/orders/domain/orders"
)

type OrdersInterface struct {
	service app.OrdersService
}

func NewOrdersInterface(service app.OrdersService) OrdersInterface {
	return OrdersInterface{service}
}

func (p OrdersInterface) MarkOrderAsPaid(orderID string) error {
	return p.service.MarkOrderAsPaid(app.MarkOrderAsPaidCommand{orders.ID(orderID)})
}
