This folder contains adapters for the interfaces (ports) defined in the orders module.

Both driven and driver adapters can be found here:

For example:
- `in_memory_order_repo.go` - implements the OrderRepository interface using a basic in-memory map
- `grpc_order_service.go` - exposes the OrderService interface as a grpc service
- `http_order_service.go` - exposes the OrderService interface as HTTP endpoints

NOTE: Different implementations of hexagonal architecture may group adapter implementations by technology, not by domain.
- For example, `gin_order_service.go` might go in the `gin` module, instead of in here in the `orders` module. It's mostly personal preference.
