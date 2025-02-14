-- +goose Up
-- +goose StatementBegin
create table threads(
    id uuid primary key,
    title text not null,
    body text not null,
    date timestamp not null default current_timestamp,
    username uuid not null,
    foreign key(username) references users(id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table threads;
-- +goose StatementEnd
