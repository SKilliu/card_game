CREATE TABLE IF NOT EXISTS users (
    id varchar(36) primary key,
    name varchar(255) not null,
    hashed_password varchar(255) not null,
    email varchar(255) unique
);

CREATE TABLE IF NOT EXISTS items (
    id varchar(36) primary key,
    name varchar(255) not null,
    image varchar(500) not null
);