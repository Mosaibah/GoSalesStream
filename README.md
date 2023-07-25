# GoSalesStream
GoSalesStream is a project that simulates E-commerce transactions and provides real-time analytics based on these transactions. The project is written in Golang, uses gRPC for communication, and stores data in a CockroachDB database.

## Prerequisites

- Git
- Docker
- Kubernetes
- Postman
- A SQL client (e.g., DBeaver)

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/Mosaibah/GoSalesStream
    ```

2. Navigate into the project directory:

    ```bash
    cd GoSalesStream
    ```

3. Create a new database named `GoSalesStream`.

4. Open the `/db_setup.sql` file, copy the SQL commands.

5. Paste and execute the SQL commands in your SQL client to set up the database schema. type (option + x) to run all commands

## Deployment

1. Open a terminal and navigate to the project directory.

2. Run the Kubernetes deployment script:

    ```bash
    ./k8s/deploy.sh
    ```

3. Forward the ports for the transaction service:

    ```bash
    kubectl port-forward -n go-sales-stream svc/transaction-service 9090:80
    ```
    
4. Forward the ports for the analytics service:

    ```bash
    kubectl port-forward -n go-sales-stream svc/analytics-service 9091:80
    ```

## Testing with Postman

1. Open Postman.

2. Add a gRPC request for the transactions service. Import the protobuf file from `GoSalesStream/Transaction.proto`.

3. Add another gRPC request for the analytics service. Import the protobuf file from `GoSalesStream/Analytics/analytics.proto`.

Now, you can simulate transactions and retrieve real-time analytics using the imported gRPC requests in Postman.

## Test Data 
### GetTransaction
```json
{
    "TransactionId": 2
}
```


### GetTransaction
```json
{
    "TransactionId": 2
}
```

### CreateTransaction
```json
{
    "transaction": 
        {
            "CustomerId": 1,
            "ProductId": 2,
            "Price": 122,
            "Quantity": 45
        }
}
```

### GetSalesByProduct
```json
{
    "ProductId": 2
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
