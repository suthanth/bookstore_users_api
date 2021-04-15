
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table user (
    id int,
    first_name varchar(100),
    last_name varchar(100),
    email varchar(100),
    date_created varchar(100),
    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table if exists user;