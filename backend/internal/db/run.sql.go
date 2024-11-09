// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: run.sql

package db

import (
	"context"
)

const getFlag = `-- name: GetFlag :one
SELECT id, question, answer, answer_option, score, seq_num, visited, correct, team_id, created_at, updated_at FROM flag WHERE id = $1
`

// GetFlag
//
//	SELECT id, question, answer, answer_option, score, seq_num, visited, correct, team_id, created_at, updated_at FROM flag WHERE id = $1
func (q *Queries) GetFlag(ctx context.Context, id string) (Flag, error) {
	row := q.db.QueryRow(ctx, getFlag, id)
	var i Flag
	err := row.Scan(
		&i.ID,
		&i.Question,
		&i.Answer,
		&i.AnswerOption,
		&i.Score,
		&i.SeqNum,
		&i.Visited,
		&i.Correct,
		&i.TeamID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTeam = `-- name: GetTeam :one
SELECT id, team_name, secret FROM team LIMIT 1
`

// GetTeam
//
//	SELECT id, team_name, secret FROM team LIMIT 1
func (q *Queries) GetTeam(ctx context.Context) (Team, error) {
	row := q.db.QueryRow(ctx, getTeam)
	var i Team
	err := row.Scan(&i.ID, &i.TeamName, &i.Secret)
	return i, err
}
