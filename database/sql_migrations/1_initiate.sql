-- +migrate Up
-- +migrate StatementBegin

create table users (
    id SERIAL PRIMARY KEY NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL
);

create table category (
    id SERIAL PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL
);

create table expenses (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    category_id INT NOT NULL,
    types INT NOT NULL,  -- 1 = income, 2 = outcome
    dates date NOT NULL,
    amount BIGINT NOT NULL,
    descriptions text,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES category(id),
    CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users(id)
);


-- +migrate StatementEnd