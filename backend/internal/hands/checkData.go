package hands

import (
	"fmt"
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func CheckValidData(hand *types.PostHandScore) error {
	err := checkAgari(hand)
	if err != nil {
		return err
	}
	err = checkMinMenzenTiles(hand)
	if err != nil {
		return err
	}
	err = checkTotalTileAmount(hand)
	if err != nil {
		return err
	}
	err = checkCallTileMultiples(hand)
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
	err = checkKaze(hand)
	if err != nil {
		return err
	}

	return nil
}

func checkMinMenzenTiles(hand *types.PostHandScore) error {
	if len(hand.Hand.Menzen) <= 1 {
		return fmt.Errorf("Must have at least 2 menzen tiles")
	}
	return nil
}

func checkCallTileMultiples(hand *types.PostHandScore) error {
	if len(hand.Hand.Chi)%3 != 0 || len(hand.Hand.Pon)%3 != 0 {
		return fmt.Errorf("Calls Chi and Pon must be multiple of 3")
	}
	if len(hand.Hand.Kan)%4 != 0 || len(hand.Hand.Ankan)%4 != 0 {
		return fmt.Errorf("Calls Kan and Ankan must be multiple of 4")
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
		//if aka count as normal
		if tile == "M5A" || tile == "S5A" || tile == "P5A" {
			tile = tile[:len(tile)-1]
		}
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

func checkKaze(hand *types.PostHandScore) error {
	var kaze = map[string]bool{
		"H1": true,
		"H2": true,
		"H3": true,
		"H4": true,
	}
	if !kaze[hand.Jikaze] || !kaze[hand.Bakaze] {
		return fmt.Errorf("Bakaze and Jikaze must be in H1,H2,H3,H4 (Ton,Nan,Sha,Pei)")
	}

	return nil
}

func checkAgari(hand *types.PostHandScore) error {
	if !hand.Ron && !hand.Tsumo {
		return fmt.Errorf("Must have agari type ron or tsumo")
	}
	if !types.Tiles[hand.Hand.Agari] {
		return fmt.Errorf("%s is not a valid agari tile", hand.Hand.Agari)

	}
	return nil
}
