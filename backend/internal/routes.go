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
	var postHandScore types.PostHandScore
	err := utils.ParseJSON(r, &postHandScore)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	hand := types.HandPartsBlocks{
		Menzen: make([][]string, 0),
		Chi:    make([][]string, 0),
		Pon:    make([][]string, 0),
		Kan:    make([][]string, 0),
		Ankan:  make([][]string, 0),
	}
	returnHandScore := types.ReturnHandScore{Hand: &hand, Yaku: []*types.YakuComponet{}, FuComponet: []*types.FuComponet{}}

	utils.WriteJSON(w, http.StatusCreated, returnHandScore)

}
