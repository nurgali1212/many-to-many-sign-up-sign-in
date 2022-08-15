BEGIN;
CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password varchar(255) not null
    
);

CREATE TABLE IF NOT EXISTS toys(
    id serial PRIMARY KEY,
    person varchar(255) NOT NULL,
    movie TEXT NOT NULL

);

CREATE TABLE IF NOT EXISTS users_toys (
    id serial PRIMARY KEY,
    users_id  int REFERENCES users(id) on delete cascade    not null,
    toys_id int REFERENCES toys(id) on delete cascade   not null
);


CREATE TABLE IF NOT EXISTS category(
    id serial PRIMARY KEY,
    genre varchar(255) NOT NULL

);





CREATE TABLE IF NOT EXISTS toys_category (
    id serial PRIMARY KEY,
    toys_id  int REFERENCES toys(id) on delete cascade    not null,
    category_id int REFERENCES category(id) on delete cascade    not null

);









COMMIT;