-- name: CreateThread :one
insert into threads (id, title, body, date, owner) values(gen_random_uuid(), @title, @body, now(), @owner) returning id, title, owner;
-- name: GetThreadById :one
select threads.id, threads.title, threads.body, threads.date, users.username from threads join users on threads.owner = users.id where threads.id = @id;
