create table sessions
(
    id            bigserial primary key,
    user_id       bigint       not null,
    refresh_token varchar(256) not null unique,
    created_at    timestamptz  not null default now(),
    updated_at    timestamptz  not null default now(),
    expired_at    timestamptz  not null default now() + interval '30 days'
);