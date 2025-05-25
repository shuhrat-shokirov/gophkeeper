create table if not exists users
(
    id         bigserial primary key,
    email      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);