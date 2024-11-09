-- name: GetTeam :one
SELECT
    *
FROM
    team
LIMIT 1;

-- name: GetFlag :one
SELECT
    *
FROM
    flag
WHERE
    id = $1;

-- name: GetTeamBySecret :one
SELECT
    *
FROM
    team
WHERE
    secret = $1;

-- name: MarkFlagAsCorrect :exec
UPDATE
    flag
SET
    correct = TRUE
WHERE
    id = $1;

-- name: MarkFlagAsIncorrect :exec
UPDATE
    flag
SET
    correct = FALSE
WHERE
    id = $1;

-- name: CheckIfAnswered :one
SELECT
    CASE WHEN correct IS NOT NULL
        AND id = $1 THEN
        TRUE
    ELSE
        FALSE
    END
FROM
    flag;

-- name: AssignQuestionToTeam :exec
WITH to_update AS (
    SELECT
        id
    FROM
        flag
    WHERE
        team_id IS NULL
    LIMIT 5)
UPDATE
    flag
SET
    team_id = $1
WHERE
    id IN (
        SELECT
            id
        FROM
            to_update);
