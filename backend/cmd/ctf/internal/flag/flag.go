package flag

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/geekodour/shd_24_ctf/backend/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type flagSubmitRequestParams struct {
	FlagID     string `json:"flag_id"`
	TeamSecret string `json:"team_secret"`
	FlagAnswer string `json:"flag_answer"`
}

type flagAttributesResponseParams struct {
	FlagID        string   `json:"flag_id"`
	Question      string   `json:"question"`
	AnswerOptions []string `json:"answer_options"`
	Score         int32    `json:"score"`
	Sequence      int32    `json:"sequence"`
}

type teamAttributesResponseParams struct {
	MapURL   string `json:"map_url"`
	TeamName string `json:"team_name"`
}

type FlagHandler struct {
	q    *db.Queries   // sqlc
	pool *pgxpool.Pool // pgx
}

func NewFlagHandler(queries *db.Queries, p *pgxpool.Pool) FlagHandler {
	return FlagHandler{
		q:    queries,
		pool: p,
	}
}

func commitOrRollback(ctx context.Context, commit bool, tx pgx.Tx) {
	var err error
	if commit {
		err = tx.Commit(ctx)
	} else {
		err = tx.Rollback(ctx)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func (h *FlagHandler) CheckResponse(ctx context.Context, params flagSubmitRequestParams) error {
	tx, err := h.pool.Begin(ctx)
	if err != nil {
		return err
	}
	qtx := h.q.WithTx(tx)
	shouldCommit := false
	defer func() {
		defer commitOrRollback(ctx, shouldCommit, tx)
	}()

	flagInfo, err := qtx.GetFlag(ctx, params.FlagID)
	if err != nil {
		return err
	}

	// check if already answered
	IsAnswered, err := qtx.CheckIfAnswered(ctx, params.FlagID)
	if err != nil {
		return fmt.Errorf("already answered: %w", err)
	}
	if IsAnswered {
		return fmt.Errorf("already answered")
	}

	// verify secret
	submittedSecret := params.TeamSecret
	_, err = qtx.GetTeamBySecret(ctx, submittedSecret)
	if err != nil {
		return fmt.Errorf("incorrect secret: %w", err)
	}
	fmt.Println(submittedSecret)

	// check if answer is correct
	submittedAnswer := params.FlagAnswer
	if submittedAnswer == flagInfo.Answer {
		err := qtx.MarkFlagAsCorrect(ctx, params.FlagID)
		if err != nil {
			return err
		}
		shouldCommit = true
	} else {
		err := qtx.MarkFlagAsIncorrect(ctx, params.FlagID)
		if err != nil {
			return err
		}
		shouldCommit = true
	}

	shouldCommit = true
	return nil
}

func (h *FlagHandler) TeamInfo(w http.ResponseWriter, r *http.Request) {
	teamSecret := chi.URLParam(r, "team_secret")
	team, err := h.q.GetTeamBySecret(r.Context(), teamSecret)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	teamAttrs := teamAttributesResponseParams{
		MapURL:   team.MapUrl,
		TeamName: team.TeamName,
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(teamAttrs)
}

func (h *FlagHandler) Get(w http.ResponseWriter, r *http.Request) {
	flagID := chi.URLParam(r, "f_id")
	flagInfo, err := h.q.GetFlag(r.Context(), flagID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	var answerOpts = []string{}
	err = json.Unmarshal(flagInfo.AnswerOption, &answerOpts)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	flagAttrs := flagAttributesResponseParams{
		FlagID:        flagID,
		Question:      flagInfo.Question,
		AnswerOptions: answerOpts,
		Score:         flagInfo.Score,
		Sequence:      flagInfo.SeqNum,
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(flagAttrs)
}

func (h *FlagHandler) Submit(w http.ResponseWriter, r *http.Request) {
	var params flagSubmitRequestParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

    // data := []byte("Hello, this is some binary data")
	err = h.CheckResponse(r.Context(), params)
	if err != nil {
		errResp := map[string]string{"error": err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResp)
		return
	}
	fmt.Println("NO ERROR REPORTED")
	w.WriteHeader(http.StatusOK)
}
