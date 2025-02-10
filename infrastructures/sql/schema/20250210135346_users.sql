-- +goose Up
-- +goose StatementBegin
create table users (
    id uuid primary key,
    username text not null,
    password text not null,
    name text not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
