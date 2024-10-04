package internal

import (
	"net/http"

	"github.com/KainoaGardner/webMahjongCalc/internal/hands"
	"github.com/KainoaGardner/webMahjongCalc/internal/tenpai"
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
	r.Post("/tenpai", h.tenpai)
	r.Post("/shanten", h.shanten)

}

func (h *Handler) scoreHand(w http.ResponseWriter, r *http.Request) {
	var postHandScore types.PostHandScore
	err := utils.ParseJSON(r, &postHandScore)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	returnHandScore, err := hands.GetHandScore(&postHandScore)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, returnHandScore)
}

func (h *Handler) tenpai(w http.ResponseWriter, r *http.Request) {
	var hand types.HandParts
	err := utils.ParseJSON(r, &hand)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	winningTiles, err := tenpai.GetWinningTiles(&hand)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, winningTiles)
}

func (h *Handler) shanten(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("shanten"))
}
