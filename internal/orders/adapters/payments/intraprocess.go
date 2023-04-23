package payments

import (
	"github.com/ttlzx/monolith-microservice-shop/internal/common/price"
	"github.com/ttlzx/monolith-microservice-shop/internal/orders/domain/orders"
	"github.com/ttlzx/monolith-microservice-shop/internal/payments/ports/intraprocess"
)

type IntraprocessService struct {
	orders chan <- intraprocess.OrderToProcess
}

func NewIntraprocessService(ordersChannel chan <- intraprocess.OrderToProcess) IntraprocessService {
	return IntraprocessService{ordersChannel}
}

func (i IntraprocessService) InitializeOrderPayment(id orders.ID, price price.Price) error {
	i.orders <- intraprocess.OrderToProcess{string(id), price}
	return nil
}
