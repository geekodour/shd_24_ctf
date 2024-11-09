-- name: GetTeam :one
SELECT * FROM team LIMIT 1;

-- name: GetFlag :one
SELECT * FROM flag WHERE id = $1;
