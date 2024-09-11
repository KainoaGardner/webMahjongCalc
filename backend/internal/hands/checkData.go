package hands

import (
	"fmt"
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func CheckValidData(hand *types.PostHandScore) error {

	err := checkTotalTileAmount(hand)
	if err != nil {
		return err
	}
	err = checkEachTileAmount(hand)
	if err != nil {
		return err
	}

	err = checkValidTiles(hand)
	if err != nil {
		return err
	}

	return nil
}

func checkTotalTileAmount(hand *types.PostHandScore) error {
	menzenLen := len(hand.Hand.Menzen)
	chiLen := len(hand.Hand.Chi)
	ponLen := len(hand.Hand.Pon)
	kanLen := len(hand.Hand.Kan)
	ankanLen := len(hand.Hand.Ankan)

	//get total tileAmount
	//subtact 1 for each kan group
	tileAmount := menzenLen + chiLen + ponLen + (kanLen - (kanLen / 4)) + (ankanLen - (ankanLen / 4))

	if tileAmount != 14 {
		return fmt.Errorf("Must have 14 tiles (kan count as 3)")
	}
	return nil
}

func getEachTileAmount(tiles []string, count map[string]int) {
	for _, tile := range tiles {
		count[tile] += 1
	}
}

func checkEachTileAmount(hand *types.PostHandScore) error {
	count := make(map[string]int)

	getEachTileAmount(hand.Hand.Menzen, count)
	getEachTileAmount(hand.Hand.Chi, count)
	getEachTileAmount(hand.Hand.Pon, count)
	getEachTileAmount(hand.Hand.Kan, count)
	getEachTileAmount(hand.Hand.Ankan, count)
	getEachTileAmount(hand.Dora, count)

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

func checkValidTiles(hand *types.PostHandScore) error {
	err := checkValidTile(hand.Hand.Menzen)
	if err != nil {
		return err
	}
	err = checkValidTile(hand.Hand.Chi)
	if err != nil {
		return err
	}
	err = checkValidTile(hand.Hand.Pon)
	if err != nil {
		return err
	}
	err = checkValidTile(hand.Hand.Kan)
	if err != nil {
		return err
	}
	err = checkValidTile(hand.Hand.Ankan)
	if err != nil {
		return err
	}
	err = checkValidTile(hand.Dora)
	if err != nil {
		return err
	}

	return nil

}
