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
insert into toys (person,movie)
values
('Spider Man','Avengers'),
('Blazz Lighter','Toys History'),
('Grooth','Guardians of Galaxy'),
('Iondu','civil war');

CREATE TABLE IF NOT EXISTS category(
    id serial PRIMARY KEY,
    genre varchar(255) NOT NULL

);
insert into category (genre)
values
('Fantasy'),
('Drama'),
('Comedy'),
('Horror');

CREATE TABLE IF NOT EXISTS userstoys (
    id serial PRIMARY KEY,
    users_id  int,
    toys_id int,
    FOREIGN KEY (users_id) REFERENCES users(id),
    FOREIGN KEY (toys_id) REFERENCES toys(id),
    UNIQUE (users_id, toys_id)
);



CREATE TABLE IF NOT EXISTS toyscategory (
    id serial PRIMARY KEY,
    toys_id  int,
    category_id int,
    FOREIGN KEY (toys_id) REFERENCES toys(id),
    FOREIGN KEY (category_id) REFERENCES category(id),
    UNIQUE (toys_id, category_id)
);






-- insert into userstoys(users_id,toys_id)
-- values
-- (1,1),
-- (2,1),
-- (1,3),
-- (2,1);




insert into toyscategory(toys_id,category_id)
values
(1,1),
(2,1),
(3,3),
(4,1);


COMMIT;