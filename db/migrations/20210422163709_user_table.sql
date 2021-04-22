
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table users (
    id int not null auto_increment,
    first_name varchar(100) not null,
    last_name varchar(100) not null,
    email varchar(100) not null,
    password varchar(1000) not null,
    created_at timestamp not null,
    updated_at timestamp DEFAULT current_timestamp,
    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table if exists users;
