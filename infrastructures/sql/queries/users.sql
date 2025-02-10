-- name: CreateUser :one
insert into users (id, username, password, name, created_at, updated_at) values(gen_random_uuid(), lower(@username), @password, lower(@name), now(), now()) returning id, username, name;

-- name: GetUsers :many
select * from users;

-- name: GetByUsername :one
select * from users where username = @username;

-- name: GetByID :one
select * from users where id = @id;
