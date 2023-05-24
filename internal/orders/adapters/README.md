This folder contains adapters for the interfaces (ports) defined in the orders module.

Both driven and driver adapters can be found here:

For example:
    - `in_memory_order_repo.go` - implements the OrderRepository interface using a basic in-memory map
    - `grpc_order_service.go` - exposes the OrderService interface as a grpc service
    - `http_order_service.go` - exposes the OrderService interface as HTTP endpoints
