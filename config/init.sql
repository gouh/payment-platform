CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE customers
(
    id    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name  VARCHAR(255)        NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE merchants
(
    id   UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

CREATE TABLE tokenized_cards
(
    token            VARCHAR(255) PRIMARY KEY,
    last_four_digits CHAR(4)     NOT NULL,
    expiry_month     INTEGER     NOT NULL,
    expiry_year      INTEGER     NOT NULL,
    card_type        VARCHAR(50) NOT NULL,
    customer_id      UUID        NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers (id)
);

CREATE TABLE payments
(
    id               UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    merchant_id      UUID                        NOT NULL,
    token            VARCHAR(255)                NOT NULL,
    amount           NUMERIC(10, 2)              NOT NULL,
    status           VARCHAR(50)                 NOT NULL,
    transaction_date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    FOREIGN KEY (merchant_id) REFERENCES merchants (id),
    FOREIGN KEY (token) REFERENCES tokenized_cards (token)
);
