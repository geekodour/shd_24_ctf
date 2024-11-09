package flag

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/geekodour/shd_24_ctf/backend/internal/db"
	"github.com/go-chi/chi/v5"
)

type flagSubmitRequestParams struct {
    FlagID      string `json:"flag_id"`
    TeamSecret  string `json:"team_secret"`
    FlagAnswer  string `json:"flag_answer"`
}
// type flagSubmitResponseParams struct {}

type flagAttributesResponseParams struct {
    FlagID          string   `json:"flag_id"`
    Question        string   `json:"question"`
    AnsswerOptions  []string `json:"answer_options"`
    Score          int      `json:"score"`
    Sequence       int      `json:"sequence"`
}

type FlagHandler struct {
	q *db.Queries // sqlc
}

func NewFlagHandler(queries *db.Queries) FlagHandler {
		return FlagHandler{
				q: queries,
		}
}

func (h *FlagHandler) Get(w http.ResponseWriter, r *http.Request) {
	flagID := chi.URLParam(r, "f_id")
	flagAttrs := flagAttributesResponseParams{
			FlagID: flagID,
	}

	// TODO: Query DB

	// err := json.NewDecoder(r.Body).Decode(&flagAttrs)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(flagAttrs)
}

func (h *FlagHandler) Submit(w http.ResponseWriter, r *http.Request) {
	var params flagSubmitRequestParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println(params)

	// TODO Check response, Query DB
	// if incorrect

	w.WriteHeader(http.StatusAccepted)
	// w.WriteHeader(http.StatusOK)
}
