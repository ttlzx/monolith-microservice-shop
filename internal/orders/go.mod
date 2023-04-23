module github.com/ttlzx/monolith-microservice-shop/internal/orders

go 1.20

require (
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/render v1.0.2
	github.com/stretchr/testify v1.8.2
)

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/streadway/amqp v1.0.0
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/ttlzx/monolith-microservice-shop/internal/common => ../common/
