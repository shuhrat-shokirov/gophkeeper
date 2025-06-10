CREATE TABLE IF NOT EXISTS binary_data
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id    BIGINT    NOT NULL REFERENCES users (id),
    title      TEXT      NOT NULL,
    content    TEXT     NOT NULL,
    note       TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);