# GoSalesStream
GoSalesStream is a project that simulates E-commerce transactions and provides real-time analytics based on these transactions. The project is written in Golang, uses gRPC for communication, and stores data in a CockroachDB database.

## Prerequisites

- Go
- CockroachDB
- A SQL client (e.g., DBeaver)
- Postman
  

## Setup

1. Setup DB
   ```bash
   cockroach start --insecure --store=node1 --listen-addr=localhost:26260 --http-addr=localhost:8089 --join=localhost:26260 --background
   ```
   then
   ```bash
   cockroach init --insecure --host=localhost:26260
   ```
2. Create database & seed data <br>
 - Open the `/db_setup.sql` file, copy the SQL commands. <br>
 - Paste and execute the SQL commands in your SQL client to set up the database schema. type (option + x) to run all commands


3. Clone the repository:

    ```bash
    git clone https://github.com/Mosaibah/GoSalesStream
    ```

4. Navigate into the project directory:

    ```bash
    cd GoSalesStream
    ```

5. Build the project
   ```bash
   task build
   ```

6. Run transaction service
    ```bash 
    task run-transaction
    ```
7. Run analytics service 
    <br> open new terminal
    ```bash
    task run-analytics
    ```

## Testing with Postman

1. Open Postman.

2. Add a gRPC request for the transactions service. Import the protobuf file from `GoSalesStream/packages/proto/transaction/v1/transaction.proto`.

3. Add another gRPC request for the analytics service. Import the protobuf file from `GoSalesStream/packages/proto/analytics/v1/analytics.proto`.

Now, you can simulate transactions and retrieve real-time analytics using the imported gRPC requests in Postman.

## Test Data 

### GetTransaction
```json
{
    "transaction_id": 2
}
```

### CreateTransaction
```json
{
    "transaction": 
        {
            "customer_id": 1,
            "product_id": 2,
            "price": 122,
            "quantity": 45
        }
}
```

### GetSalesByProduct
```json
{
    "product_id": 2
}
```
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
