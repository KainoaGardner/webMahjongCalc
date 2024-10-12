package hands

import (
	"fmt"
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func checkValidData(hand *types.PostHandScore) error {
	err := checkAgari(hand)
	if err != nil {
		return err
	}
	err = checkMinMenzenTiles(hand.Hand)
	if err != nil {
		return err
	}
	err = checkTotalTileAmount(hand.Hand)
	if err != nil {
		return err
	}
	err = checkEachTileAmount(hand)
	if err != nil {
		return err
	}
	err = CheckCallTileMultiples(hand.Hand)
	if err != nil {
		return err
	}
	err = checkValidTiles(hand)
	if err != nil {
		return err
	}
	err = checkKaze(hand.ScoringParts)
	if err != nil {
		return err
	}
	err = checkMultipleAkaDora(hand)
	if err != nil {
		return err
	}
	err = checkOpenRiichi(hand)
	if err != nil {
		return err
	}

	return nil
}

func checkMinMenzenTiles(hand *types.HandParts) error {
	if len(hand.Menzen) < 1 {
		return fmt.Errorf("Must have at least 1 menzen tile")
	}
	return nil
}

func CheckCallTileMultiples(hand *types.HandParts) error {
	if len(hand.Chi)%3 != 0 || len(hand.Pon)%3 != 0 {
		return fmt.Errorf("Calls Chi and Pon must be multiple of 3")
	}
	if len(hand.Kan)%4 != 0 || len(hand.Ankan)%4 != 0 {
		return fmt.Errorf("Calls Kan and Ankan must be multiple of 4")
	}

	return nil
}

func checkTotalTileAmount(hand *types.HandParts) error {
	menzenLen := len(hand.Menzen)
	chiLen := len(hand.Chi)
	ponLen := len(hand.Pon)
	kanLen := len(hand.Kan)
	ankanLen := len(hand.Ankan)

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
		count[tile[:2]] += 1
	}
}

func checkEachTileAmount(hand *types.PostHandScore) error {
	count := make(map[string]int)

	getEachTileAmount(hand.Hand.Menzen, count)
	getEachTileAmount(hand.Hand.Chi, count)
	getEachTileAmount(hand.Hand.Pon, count)
	getEachTileAmount(hand.Hand.Kan, count)
	getEachTileAmount(hand.Hand.Ankan, count)
	getEachTileAmount(getDoraIndicator(hand.ScoringParts.Dora), count)
	getEachTileAmount(getDoraIndicator(hand.ScoringParts.Uradora), count)

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
	err = checkValidTile(getDoraIndicator(hand.ScoringParts.Dora))
	if err != nil {
		return err
	}
	err = checkValidTile(getDoraIndicator(hand.ScoringParts.Uradora))
	if err != nil {
		return err
	}

	return nil

}

func checkKaze(hand *types.HandScoringParts) error {
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
	if !hand.ScoringParts.Ron && !hand.ScoringParts.Tsumo {
		return fmt.Errorf("Must have agari type ron or tsumo")
	}
	if hand.ScoringParts.Ron && hand.ScoringParts.Tsumo {
		return fmt.Errorf("Must only one agari type ron or tsumo")
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

func checkMultipleAkaDora(hand *types.PostHandScore) error {
	count := make(map[string]int)

	getEachTileAkaDora(hand.Hand.Menzen, count)
	getEachTileAkaDora(hand.Hand.Chi, count)
	getEachTileAkaDora(hand.Hand.Pon, count)
	getEachTileAkaDora(hand.Hand.Kan, count)
	getEachTileAkaDora(hand.Hand.Ankan, count)
	getEachTileAkaDora(getDoraIndicator(hand.ScoringParts.Dora), count)
	getEachTileAkaDora(getDoraIndicator(hand.ScoringParts.Uradora), count)

	for _, amount := range count {
		if amount > 1 {
			return fmt.Errorf("Can only have at most 1 of each akadora")
		}

	}
	return nil
}

func checkOpenRiichi(hand *types.PostHandScore) error {
	var open bool

	if len(hand.Hand.Chi) > 0 || len(hand.Hand.Pon) > 0 || len(hand.Hand.Kan) > 0 {
		open = true
	}

	if (open && hand.ScoringParts.Riichi) || (open && hand.ScoringParts.Wriichi) {
		return fmt.Errorf("Cannot call riichi with an open hand")
	}

	return nil
}

func getDoraIndicator(dora []string) []string {
	var doraIndicator []string

	for _, doraTile := range dora {
		tileNumber := int(doraTile[1] - '0')
		tileNumber--
		if doraTile[0] == 'H' {
			if tileNumber == 0 {
				tileNumber = 4
			} else if tileNumber == 4 {
				tileNumber = 7
			}
		} else {
			if tileNumber <= 0 {
				tileNumber = 9
			}
		}

		var newTile string
		if len(doraTile) >= 3 {

			newTile = fmt.Sprintf("%c%d%c", doraTile[0], tileNumber, doraTile[2])
		} else {
			newTile = fmt.Sprintf("%c%d", doraTile[0], tileNumber)

		}
		doraIndicator = append(doraIndicator, newTile)
	}

	return doraIndicator
}
