package internal

import (
	"net/http"

	"github.com/KainoaGardner/webMahjongCalc/types"
	"github.com/KainoaGardner/webMahjongCalc/utils"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/points", h.scoreHand)

}

func (h *Handler) scoreHand(w http.ResponseWriter, r *http.Request) {
	var payload types.PostHandScore
	err := utils.ParseJSON(r, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, payload)

}
