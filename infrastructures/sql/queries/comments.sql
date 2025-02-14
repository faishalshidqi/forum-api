-- name: CreateComment :one
insert into comments(id, owner, thread, content, date) values(gen_random_uuid(), @owner, @thread, @content, now()) returning id, content, owner;
