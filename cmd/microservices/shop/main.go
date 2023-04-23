package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/ttlzx/monolish-microservice-shop/internal//shop"
	"github.com/ttlzx/monolith-microservice-shop/internal/common/cmd"
	shop_app "github.com/ttlzx/monolith-microservice-shop/internal/shop/application"
	shop_infra_product "github.com/ttlzx/monolith-microservice-shop/internal/shop/infrastructure/products"
	shop_interfaces_private_http "github.com/ttlzx/monolith-microservice-shop/internal/shop/interfaces/private/http"
	shop_interfaces_public_http "github.com/ttlzx/monolith-microservice-shop/internal/shop/interfaces/public/http"
)

func main() {
	log.Println("Starting shop microservice")

	ctx := cmd.Context()

	r := createShopMicroservice()
	server := &http.Server{Addr: os.Getenv("SHOP_SHOP_SERVICE_BIND_ADDR"), Handler: r}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()
	log.Println("Closing shop microservice")

	if err := server.Close(); err != nil {
		panic(err)
	}
}

func createShopMicroservice() *chi.Mux {
	shopProductRepo := shop_infra_product.NewMemoryRepository()
	shopProductsService := shop_app.NewProductsService(shopProductRepo, shopProductRepo)

	if err := shop.LoadShopFixtures(shopProductsService); err != nil {
		panic(err)
	}

	r := cmd.CreateRouter()

	shop_interfaces_public_http.AddRoutes(r, shopProductRepo)
	shop_interfaces_private_http.AddRoutes(r, shopProductRepo)

	return r
}
