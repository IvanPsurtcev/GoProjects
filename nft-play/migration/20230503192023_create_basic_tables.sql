-- +goose Up
-- +goose StatementBegin
create table "user"(
    id serial primary key,
    address varchar(128) UNIQUE
);

create table request(
    id serial primary key,
    applicant_addr varchar(128),
    receiver_addr varchar(128),
    status int not null default  1-- 1 - pending, 2 - accepted, 0 - declined
);

create table game(
    id serial primary key,
    applicant_addr varchar(128),
    receiver_addr varchar(128),
    winner_addr varchar(128) not null default '-',
    status int not null default 1, -- 1 - active, 2 - pending, 3 - done, 0 - canceled
    applicant_ready boolean not null default false,
    receiver_ready boolean not null default false,
    create_date timestamp not null default now(),
    applicant_nft varchar(256) not null default '',
    receiver_nft varchar(256) not null default ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "user";
drop table request;
drop table game;
-- +goose StatementEnd
