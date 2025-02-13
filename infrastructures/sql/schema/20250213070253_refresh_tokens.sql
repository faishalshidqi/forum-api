-- +goose Up
-- +goose StatementBegin
create table refresh_tokens(
    token text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table refresh_tokens;
-- +goose StatementEnd
