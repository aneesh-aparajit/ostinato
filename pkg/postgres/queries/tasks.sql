-- name: GetAllTasks :many
SELECT *
FROM tasks;

-- name: CreateTask :one
insert into tasks (id, name, execution_js, created_at, updated_at)
values ($1, $2, $3, $4, $5)
returning *;

-- name: GetTaskById :one
select *
from tasks
where id = $1;

-- name: DeleteById :exec
delete
from tasks
where id = $1;

-- name: UpdateTask :one
update tasks
set name         = $2,
    execution_js = $3
where id = $1
returning *;