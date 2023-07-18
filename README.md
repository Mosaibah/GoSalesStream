# GoSalesStream

## System Requirements:

- E-commerce Transactions Microservice: Build a microservice in Golang that simulates E-commerce transactions. Each transaction should have a unique id, a timestamp, a customer id, a -product id, a quantity, and a total price.

- gRPC Service Definitions: Use Protobuf to define gRPC services for creating, retrieving, and streaming transactions. CreateTransaction service will allow you to simulate a new transaction. GetTransaction service will return a specific transaction based on its id. StreamTransactions service should stream transactions in real-time.

- CockroachDB Integration: Store transactions in CockroachDB. There should be a function for writing transactions to the database, as well as for reading transactions.

- Temporal Workflows: Implement a Temporal workflow for processing each transaction. The workflow should handle retries in case of failures during the transaction process.

- Analytics Microservice: Implement another microservice in Golang that calculates and provides real-time analytics based on transactions. This microservice should expose gRPC services for streaming the current total sales, sales by product, and the top 5 customers.

- Unit Tests: Write unit tests for both microservices. The tests should cover gRPC handlers, Temporal workflows, and database interactions.

- Docker and Kubernetes: Both microservices should be containerized using Docker and deployed on a Kubernetes cluster.

## Functional Requirements:

- CreateTransaction: Ability to simulate a new transaction.

- GetTransaction: Ability to fetch a specific transaction based on its id.

- StreamTransactions: Ability to stream transactions in real-time.

- GetTotalSales: Ability to get the current total sales.

- GetSalesByProduct: Ability to get the current sales by product.

- GetTopCustomers: Ability to get the top 5 customers based on sales.
