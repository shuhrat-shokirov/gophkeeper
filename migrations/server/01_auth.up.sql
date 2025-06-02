create table if not exists users
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);