-- name: AddToken :one
insert into refresh_tokens (token, owner) values(@token, @owner) returning token, owner;

-- name: GetToken :one
select * from refresh_tokens where token = @token;

-- name: DeleteToken :one
delete from refresh_tokens where token = @token returning token, owner;
