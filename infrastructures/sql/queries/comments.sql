-- name: CreateComment :one
insert into comments(id, owner, thread, content, date) values(gen_random_uuid(), @owner, @thread, @content, now()) returning id, content, owner;
-- name: GetCommentsByThread :many
select comments.id, users.username, comments.date, comments.content from comments join users on comments.owner = users.id where thread = @thread;
