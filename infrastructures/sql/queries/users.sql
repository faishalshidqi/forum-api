-- name: CreateUser :one
insert into users (id, username, password, fullname, created_at, updated_at) values(gen_random_uuid(), lower(@username), @password, lower(@fullname), now(), now()) returning id, username, fullname;

-- name: GetUsers :many
select * from users;

-- name: GetByUsername :one
select * from users where username = @username;

-- name: GetUserByID :one
select * from users where id = @id;
