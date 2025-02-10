-- name: CreateUser :one
insert into users (id, username, password, name, created_at, updated_at) values(gen_random_uuid(), $1, $2, $3, now(), now()) returning id, username, name;
