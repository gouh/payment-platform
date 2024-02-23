# Payment Platform

This project is a backend application developed in Go, using the Gin framework for handling HTTP requests and `pgx` for interacting with a PostgreSQL database. The application serves as a payment platform, allowing merchants to process transactions, manage payments, and perform queries on previous payment details.

## Features

- **RESTful API**: The application exposes a RESTful API for processing payments, querying payment details, and managing merchants.
- **Payment Management**: Merchants can process payments and query details of previous payments.
- **Card Tokenization**: Implements tokenization of credit card information to enhance security.
- **Bank Simulation**: Includes a bank simulator to simulate interactions with acquiring banks. A 40% approval percentage has been configured for simulated transactions, meaning approximately 40% of payments processed through the simulator will be approved, while the rest will be considered failures.
- **Pagination and Filtering**: Supports pagination and filters in payment queries for better data management.
- **Custom Middleware**: Uses custom middleware for error handling and JSON format logs. This includes a recovery middleware that captures and logs uncaught errors in the desired format, and middleware to customize error responses to ensure they are always returned in JSON format, even when data binding errors occur.
- **Payment Authorization**: The payment authorization logic simulates communication with bank authorization systems, using a probabilistic approach to determine the approval or rejection of payment transactions based on the configured percentage.

## Technologies Used

- **Go (Golang)**: Programming language for backend development.
- **Gin**: Web framework used to create the RESTful API.
- **pgx and pgxpool**: Libraries for connection and operations with PostgreSQL.
- **goqu**: SQL query builder library for Go, used to build dynamic queries.
- **logrus**: Logging library for Go, configured to print logs in JSON format.
- **PostgreSQL**: Relational database management system used to store merchant and payment data.

## Db Model

![payment-platform](https://github.com/gouh/payment-platform/assets/13145599/1c05ae82-db23-430b-86d6-083a48146eda)


## Project Structure

```
.
├── cmd
│   └── server
│       └── main.go
├── config
│   ├── auth.go
│   ├── config.go
│   ├── database_config.go
│   └── init.sql
├── docker
│   └── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── container
│   │   ├── container.go
│   │   ├── database_dialect.go
│   │   └── database.go
│   ├── dao
│   │   ├── customer_dao.go
│   │   ├── merchant_dao.go
│   │   ├── payment_dao.go
│   │   └── tokenized_cards_dao.go
│   ├── dto
│   ├── http
│   │   ├── handlers
│   │   │   ├── auth_handler.go
│   │   │   ├── customer_handler.go
│   │   │   ├── health_handler.go
│   │   │   ├── merchant_handler.go
│   │   │   ├── payment_handler.go
│   │   │   └── tokenized_card_handler.go
│   │   ├── middleware
│   │   │   ├── auth_middleware.go
│   │   │   ├── cors_middleware.go
│   │   │   ├── error_middleware.go
│   │   │   └── recovery_middleware.go
│   │   ├── routes
│   │   │   ├── auth_routes.go
│   │   │   ├── customer_routes.go
│   │   │   ├── health_routes.go
│   │   │   ├── merchant_routes.go
│   │   │   ├── payment_routes.go
│   │   │   └── tokenized_card_routes.go
│   │   └── routes.go
│   ├── models
│   │   ├── customer.go
│   │   ├── merchant.go
│   │   ├── payment.go
│   │   └── tokenized_card.go
│   ├── requests
│   │   ├── auth_request.go
│   │   ├── customer_request.go
│   │   ├── merchant_request.go
│   │   ├── pagination_request.go
│   │   ├── payment_request.go
│   │   └── tokenized_card_request.go
│   ├── responses
│   │   ├── common_response.go
│   │   ├── customer_response.go
│   │   ├── health_response.go
│   │   ├── merchant_response.go
│   │   ├── payment_response.go
│   │   └── tokenized_card_response.go
│   ├── services
│   └── utils
│       ├── bank_simulator.go
│       ├── body_validator.go
│       └── card_tokenizer.go
├── LICENSE
└── README.md

```

## Accessing the Project

Once the containers

 are up and running, you can access the API through:

| Method | Endpoint                           | Description                                       | Authentication Required |
|--------|------------------------------------|---------------------------------------------------|-------------------------|
| POST   | `/auth/login`                      | Logs in and returns an access token.              | No                      |
| GET    | `/v1/health`                       | Checks the health status of the service.          | Yes                     |
| GET    | `/v1/merchants`                    | Retrieves a list of merchants.                    | Yes                     |
| GET    | `/v1/merchants/:id`                | Retrieves a merchant by their ID.                 | Yes                     |
| POST   | `/v1/merchants`                    | Creates a new merchant.                           | Yes                     |
| PATCH  | `/v1/merchants/:id`                | Updates a merchant's details by their ID.         | Yes                     |
| DELETE | `/v1/merchants/:id`                | Deletes a merchant by their ID.                   | Yes                     |
| GET    | `/v1/customers`                    | Retrieves a list of customers.                    | Yes                     |
| GET    | `/v1/customers/:id`                | Retrieves a customer by their ID.                 | Yes                     |
| POST   | `/v1/customers`                    | Creates a new customer.                           | Yes                     |
| PATCH  | `/v1/customers/:id`                | Updates a customer's details by their ID.         | Yes                     |
| DELETE | `/v1/customers/:id`                | Deletes a customer by their ID.                   | Yes                     |
| GET    | `/v1/customers/:id/cards`          | Retrieves tokenized cards of a customer.          | Yes                     |
| GET    | `/v1/customers/:id/cards/:token`   | Retrieves a tokenized card by its token.          | Yes                     |
| POST   | `/v1/customers/:id/cards`          | Creates a new tokenized card for a customer.      | Yes                     |
| DELETE | `/v1/customers/:id/cards/:token`   | Deletes a tokenized card by its token.            | Yes                     |
| GET    | `/v1/payments`                     | Retrieves a list of payments.                     | Yes                     |
| GET    | `/v1/payments/:id`                 | Retrieves a payment by its ID.                    | Yes                     |
| POST   | `/v1/payments`                     | Processes a new payment.                          | Yes                     |

## Contributions

Contributions are welcome. If you have any suggestions for improving the application, please consider submitting a pull request or opening an issue in the repository.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Developed with ❤ by Hugo.
