-- +goose Up
-- +goose StatementBegin
create table refresh_tokens(
    token text not null,
    owner uuid not null,
    foreign key(owner) references users(id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table refresh_tokens;
-- +goose StatementEnd
