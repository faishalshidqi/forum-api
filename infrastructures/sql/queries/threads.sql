-- name: CreateThread :one
insert into threads (id, title, body, date, owner) values(gen_random_uuid(), @title, @body, now(), @owner) returning id, title, owner;
