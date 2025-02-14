-- +goose Up
-- +goose StatementBegin
create table threads(
    id uuid primary key,
    title text not null,
    body text not null,
    date timestamp not null default current_timestamp,
    owner uuid not null,
    foreign key(owner) references users(id) on update cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table threads;
-- +goose StatementEnd
