# hexagonal-architecture-study


## Background - Hexagonal Architecture
Hexagonal Architecture is a pattern used to separate your application's business logic from its infrastructure.

In hexagonal architecture, each module contains the following:
- Core: business logic, no dependenices
- Driver ports:  interfaces used by external systems to interact with the core module
- Driven ports: interfaces used by the core module to interact with external systems

And the following are implemented outside of the module:
- Driver adapter: implementations of driver ports, using a specific framework or protocol (i.e. HTTP controllers)
- Driven adapters: implementations of driven ports, using a sepcific framework or protocol (i.e. SQL database)


Code structure is based on the examples found in the book "Event Driven Architecture in Golang" - by Michael Stack
https://github.com/PacktPublishing/Event-Driven-Architecture-in-Golang/ 

## Scenario
So this example is non-trivial, we'll implement a basic order-fufillment application here. 
Orders will be received, shipped, and delivered, with this service tracking all of those actions.


## Productionalization
To make this app more production-ready, it will need the following capabilities:
   - [x] Unit test suite
   - [x] E2E test suite
   - [x] Local ephemeral env for E2E Testing (i.e. docker-compose)
   - [ ] Persistent storage / DB 
   - [ ] CICD
   - [ ] OpenTelemetry 
        - otel loggging
        - otel spans/traces for each entrypoint (i.e. http endpoint), and each external service call
        - otel metrics (cpu, mem, net in/out, num requests)