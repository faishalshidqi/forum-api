-- +goose Up
-- +goose StatementBegin
create table comments(
    id uuid primary key,
    owner uuid not null,
    foreign key(owner) references users(id) on update cascade,
    thread uuid not null,
    foreign key(thread) references threads(id) on update cascade,
    content text not null,
    date timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table comments;
-- +goose StatementEnd
