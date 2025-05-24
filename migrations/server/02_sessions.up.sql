create table sessions
(
    id            text primary key,
    user_id       text        not null,
    refresh_token text        not null unique,
    created_at    timestamptz not null default now(),
    updated_at    timestamptz not null default now(),
    expired_at    timestamptz not null default now() + interval '30 days'
);