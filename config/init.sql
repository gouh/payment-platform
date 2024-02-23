CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE customers
(
    id    UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name  VARCHAR(255)                    NOT NULL,
    email VARCHAR(255)                    NOT NULL UNIQUE
);

CREATE TABLE merchants
(
    id   UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR(255)                    NOT NULL
);

CREATE TABLE tokenized_cards
(
    token            VARCHAR(255) NOT NULL PRIMARY KEY,
    last_four_digits CHAR(4)      NOT NULL,
    expiry_month     INTEGER      NOT NULL,
    expiry_year      INTEGER      NOT NULL,
    card_type        VARCHAR(50)  NOT NULL,
    customer_id      UUID         NOT NULL REFERENCES customers
);

CREATE TABLE payments
(
    id               UUID DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    merchant_id      UUID                            NOT NULL REFERENCES merchants,
    token            VARCHAR(255)                    NOT NULL REFERENCES tokenized_cards,
    amount           NUMERIC(10, 2)                  NOT NULL,
    status           VARCHAR(50)                     NOT NULL,
    transaction_date TIMESTAMP                       NOT NULL
);