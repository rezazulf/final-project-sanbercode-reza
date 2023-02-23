-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(256),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TYPE roles AS ENUM ('Admin', 'Customer');

CREATE TABLE users(
    id SERIAL NOT NULL PRIMARY KEY,
    username VARCHAR(256),
    password VARCHAR(256),
    balance INT NOT NULL,
    role roles,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE status(
    id SERIAL NOT NULL PRIMARY KEY,
    status VARCHAR(256)
);

CREATE TABLE product(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(256),
    description VARCHAR(256),
    price INT,
    image_url VARCHAR(256),
    stock INT NOT NULL,
    status_id BIGINT,
        foreign key (status_id) references status(id),
    category_id BIGINT,
        foreign key (category_id) references category(id),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE transactions(
    id SERIAL NOT NULL PRIMARY KEY,
    sum_item INT NOT NULL,
    payment_bills INT NOT NULL,
    product_id BIGINT NOT NULL,
        foreign key (product_id) references product(id),
    customer_id BIGINT NOT NULL,
        foreign key (customer_id) references users(id),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);


INSERT INTO users (username, password, balance, role) VALUES ('admin', '$2a$12$Q2BGBD3OmgSb09K7ehCzVenXdJpMdQnURDqWIzjVcvDgeZ8hm4S.G', 100000000, 'Admin')

-- +migrate StatementEnd
