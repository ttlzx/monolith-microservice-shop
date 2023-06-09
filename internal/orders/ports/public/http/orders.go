package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	common_http "github.com/ttlzx/monolith-microservice-shop/internal/common/http"
	"github.com/ttlzx/monolith-microservice-shop/internal/orders/app"
	"github.com/ttlzx/monolith-microservice-shop/internal/orders/domain/orders"
)

func AddRoutes(router *chi.Mux, service app.OrdersService, repository orders.Repository) {
	resource := ordersResource{service, repository}
	router.Post("/orders", resource.Post)
	router.Get("/orders/{id}/paid", resource.GetPaid)
}

type ordersResource struct {
	service app.OrdersService

	repository orders.Repository
}

func (o ordersResource) Post(w http.ResponseWriter, r *http.Request) {
	req := PostOrderRequest{}
	if err := render.Decode(r, &req); err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}

	u, _ := uuid.NewV1()
	cmd := app.PlaceOrderCommand{
		OrderID:   orders.ID(u.String()),
		ProductID: req.ProductID,
		Address:   app.PlaceOrderCommandAddress(req.Address),
	}
	if err := o.service.PlaceOrder(cmd); err != nil {
		_ = render.Render(w, r, common_http.Errinternal(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, PostOrdersResponse{
		OrderID: string(cmd.OrderID),
	})
}

type PostOrderAddress struct {
	Name     string `json:"name"`
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}

type PostOrderRequest struct {
	ProductID orders.ProductID `json:"product_id"`
	Address   PostOrderAddress `json:"address"`
}

type PostOrdersResponse struct {
	OrderID string
}

type OrderPaidView struct {
	ID     string `json:"id"`
	IsPaid bool   `json:"is_paid"`
}

func (o ordersResource) GetPaid(w http.ResponseWriter, r *http.Request) {
	order, err := o.repository.ByID(orders.ID(chi.URLParam(r, "id")))
	if err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}

	render.Respond(w, r, OrderPaidView{string(order.ID()), order.Paid()})
}
