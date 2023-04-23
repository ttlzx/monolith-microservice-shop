module github.com/ttlzx/monolith-microservice-shop/internal/payments

go 1.20

replace github.com/ttlzx/monolith-microservice-shop/internal/common => ../common/

require (
	github.com/pkg/errors v0.9.1
	github.com/streadway/amqp v1.0.0
)
