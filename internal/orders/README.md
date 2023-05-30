This folder contains the domain logic for the orders module.

Following hexagonal architecture, the domain logic is separated from any external code via interfaces. Domain code and ports exist in the root of this folder, and has no external dependencies
Both driving and driven adapters are implemented in the "adapters" folder.

### Examples

For example, `order_service.go` defines the functions exposed by the order module in an interface called `OrderService` (the driving port). 

This interface is implemented in the file `adapters/gin_order_service.go`, which maps the individual OrderService functions to HTTP endpoints using the gin framework. 
You could imagine other driving adapters for OrderService using other HTTP frameworks, or another protocol like GRPC or RabbitMQ.

Similarly, for the driven adapter, `order_repo.go` defines an interface `OrderRepo` for entity persistence. 
This interface can be implemented using various technologies: an in-memory map (`./adapters/in_memory_order_repo.go`), a postgres database, or any other storage mechanism.



### Composition
Because every module only relies on interfaces, we need somewhere in code where instances are instantiated and supplied. This all happens in our composition root, `cmd/app/main.go`. As the application gets sufficiently complex, we may end up using a Dependency-Injection framework to help this along, but for now manual dependency injection is fine. 