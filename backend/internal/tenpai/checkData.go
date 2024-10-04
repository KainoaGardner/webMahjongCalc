package tenpai

import (
	"fmt"

	"github.com/KainoaGardner/webMahjongCalc/internal/hands"
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func checkValidData(hand *types.HandParts) error {
	err := checkTotalTileAmount(hand)
	if err != nil {
		return err
	}
	err = checkEachTileAmount(hand)
	if err != nil {
		return err
	}
	err = hands.CheckCallTileMultiples(hand)
	if err != nil {
		return err
	}
	err = checkValidTiles(hand)
	if err != nil {
		return err
	}
	err = checkMultipleAkaDora(hand)
	if err != nil {
		return err
	}

	return nil
}

func checkTotalTileAmount(hand *types.HandParts) error {
	menzenLen := len(hand.Menzen)
	chiLen := len(hand.Chi)
	ponLen := len(hand.Pon)
	kanLen := len(hand.Kan)
	ankanLen := len(hand.Ankan)

	tileAmount := menzenLen + chiLen + ponLen + (kanLen - (kanLen / 4)) + (ankanLen - (ankanLen / 4))

	if tileAmount != 13 {
		return fmt.Errorf("Must have 13 tiles (kan count as 3)")
	}
	return nil

}

func getEachTileAmount(tiles []string, count map[string]int) {
	for _, tile := range tiles {
		//if aka count as normal
		count[tile[:2]] += 1
	}
}

func checkEachTileAmount(hand *types.HandParts) error {
	count := make(map[string]int)

	getEachTileAmount(hand.Menzen, count)
	getEachTileAmount(hand.Chi, count)
	getEachTileAmount(hand.Pon, count)
	getEachTileAmount(hand.Kan, count)
	getEachTileAmount(hand.Ankan, count)

	for _, total := range count {
		if total > 4 {
			return fmt.Errorf("Max of 4 tiles for each tile")
		}
	}

	return nil
}
func checkValidTile(tiles []string) error {
	for _, tile := range tiles {
		if !types.Tiles[tile] {
			return fmt.Errorf("%s is not a valid tile", tile)
		}
	}
	return nil
}

func checkValidTiles(hand *types.HandParts) error {
	err := checkValidTile(hand.Menzen)
	if err != nil {
		return err
	}
	err = checkValidTile(hand.Chi)
	if err != nil {
		return err
	}
	err = checkValidTile(hand.Pon)
	if err != nil {
		return err
	}
	err = checkValidTile(hand.Kan)
	if err != nil {
		return err
	}
	err = checkValidTile(hand.Ankan)
	if err != nil {
		return err
	}

	return nil

}

func getEachTileAkaDora(tiles []string, count map[string]int) {
	for _, tile := range tiles {
		if tile == "M5A" || tile == "S5A" || tile == "P5A" {
			count[tile] += 1
		}
	}
}

func checkMultipleAkaDora(hand *types.HandParts) error {
	count := make(map[string]int)

	getEachTileAkaDora(hand.Menzen, count)
	getEachTileAkaDora(hand.Chi, count)
	getEachTileAkaDora(hand.Pon, count)
	getEachTileAkaDora(hand.Kan, count)
	getEachTileAkaDora(hand.Ankan, count)

	for _, amount := range count {
		if amount > 1 {
			return fmt.Errorf("Can only have at most 1 of each akadora")
		}

	}
	return nil
}
