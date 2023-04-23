package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/ttlzx/monolith-microservice-shop/internal/common/cmd"
	orders_app "github.com/ttlzx/monolith-microservice-shop/internal/orders/application"
	orders_infra_orders "github.com/ttlzx/monolith-microservice-shop/internal/orders/infrastructure/orders"
	orders_infra_payments "github.com/ttlzx/monolith-microservice-shop/internal/orders/infrastructure/payments"
	orders_infra_product "github.com/ttlzx/monolith-microservice-shop/internal/orders/infrastructure/shop"
	orders_private_http "github.com/ttlzx/monolith-microservice-shop/internal/orders/interfaces/private/http"
	orders_public_http "github.com/ttlzx/monolith-microservice-shop/internal/orders/interfaces/public/http"
)

func main() {
	log.Println("Starting orders microservice")

	ctx := cmd.Context()

	r, closeFn := createOrdersMicroservice()
	defer closeFn()

	server := &http.Server{Addr: os.Getenv("SHOP_ORDERS_SERVICE_BIND_ADDR"), Handler: r}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()
	log.Println("Closing orders microservice")

	if err := server.Close(); err != nil {
		panic(err)
	}
}

func createOrdersMicroservice() (router *chi.Mux, closeFn func()) {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))

	shopHTTPClient := orders_infra_product.NewHTTPClient(os.Getenv("SHOP_SHOP_SERVICE_ADDR"))

	ordersToPayQueue, err := orders_infra_payments.NewAMQPService(
		fmt.Sprintf("amqp://%s/", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_RABBITMQ_ORDERS_TO_PAY_QUEUE"),
	)
	if err != nil {
		panic(err)
	}

	ordersRepo := orders_infra_orders.NewMemoryRepository()
	ordersService := orders_app.NewOrdersService(
		shopHTTPClient,
		ordersToPayQueue,
		ordersRepo,
	)

	r := cmd.CreateRouter()

	orders_public_http.AddRoutes(r, ordersService, ordersRepo)
	orders_private_http.AddRoutes(r, ordersService, ordersRepo)

	return r, func() {
		err := ordersToPayQueue.Close()
		if err != nil {
			log.Printf("cannot close orders queue: %s", err)
		}
	}
}
