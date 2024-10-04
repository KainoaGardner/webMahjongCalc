package tenpai

import (
	"github.com/KainoaGardner/webMahjongCalc/internal/hands"
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func GetWinningTiles(hand *types.HandParts) ([]string, error) {
	err := checkValidData(hand)
	if err != nil {
		return nil, err
	}
	var winningTiles []string

	for tile := range types.Tiles {
		//skip akadora
		if len(tile) == 3 {
			continue
		}

		hand.Menzen = append(hand.Menzen, tile)
		validHands, err := hands.GetValidHands(hand)
		if err == nil && len(validHands) > 0 {
			winningTiles = append(winningTiles, tile)
		}
		hand.Menzen = hand.Menzen[:len(hand.Menzen)-1]

	}

	return winningTiles, nil
}
