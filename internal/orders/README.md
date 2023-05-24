This folder contains the domain logic for the orders module.

Following hexagonal architecture, this module should not depend on any external implementations.

We define ports (interfaces) for interaction with external service, and those ports are implemented in the "adapters" folder.
