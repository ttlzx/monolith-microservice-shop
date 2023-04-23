package http

import (
	"net/http"

	"github.com/ThreeDotsLabs/monolith-microservice-shop/pkg/orders/app"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	common_http "github.com/ttlzx/monolith-microservice-shop/internal/common/http"
	"github.com/ttlzx/monolith-microservice-shop/internal/orders/domain/orders"
)

func AddRoutes(router *chi.Mux, service app.OrdersService, repository orders.Repository) {
	resource := ordersResource{service, repository}
	router.Post("/orders/{id}/paid", resource.PostPaid)
}

type ordersResource struct {
	service app.OrdersService

	repository orders.Repository
}

func (o ordersResource) PostPaid(w http.ResponseWriter, r *http.Request) {
	cmd := app.MarkOrderAsPaidCommand{
		OrderID: orders.ID(chi.URLParam(r, "id")),
	}

	if err := o.service.MarkOrderAsPaid(cmd); err != nil {
		_ = render.Render(w, r, common_http.Errinternal(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
