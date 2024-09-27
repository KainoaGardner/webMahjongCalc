package hands

import (
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func getYakuman(hand *types.WinningHand) error {
	tenhou(hand)
	chiihou(hand)
	suuankou(hand)
	kokushimusou(hand)

	daisangen(hand)
	tsuuiisou(hand)
	chinroutou(hand)
	ryuuiisou(hand)
	suukantsu(hand)

	//not done
	chuurenpoutou(hand)
	shousuushi(hand)
	daisuushi(hand)

	return nil

}

func tenhou(hand *types.WinningHand) {
	if hand.ScoringParts.Tenhou && !checkOpenHand(hand) {
		tenhou := types.YakumanComponet{Yakuman: 1, Title: "Tenhou"}
		hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &tenhou)
	}

}

func chiihou(hand *types.WinningHand) {
	if hand.ScoringParts.Tenhou && !checkOpenHand(hand) {
		chiihou := types.YakumanComponet{Yakuman: 1, Title: "Chiihou"}
		hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &chiihou)
	}

}

func suuankou(hand *types.WinningHand) {
	koutsuCount := getMenzenKoutsuCount(hand.HandParts.Menzen) + len(hand.HandParts.Ankan)
	if koutsuCount >= 4 {

		head := getHead(hand.HandParts.Menzen)

		//if tanki pattern
		if head != nil && hand.HandParts.Agari[:2] == head[0][:2] {
			suuankouTanki := types.YakumanComponet{Yakuman: 2, Title: "Suuankou Tanki"}
			hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &suuankouTanki)
		} else {

			if hand.ScoringParts.Tsumo {
				suuankou := types.YakumanComponet{Yakuman: 1, Title: "Suuankou"}
				hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &suuankou)

			}
		}
	}
}

func kokushimusou(hand *types.WinningHand) {
	if !checkOpenHand(hand) {
		unsortedHand := getUnsortedHand(hand)
		handCount := getTileMap(unsortedHand)

		if checkKokushiAllTiles(handCount) {
			if checkKokushiThirteenWait(hand.HandParts) {
				kokushiJuusanmen := types.YakumanComponet{Yakuman: 2, Title: "Kokushimusou Juusanmenmachi"}
				hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &kokushiJuusanmen)
			} else {
				kokushi := types.YakumanComponet{Yakuman: 1, Title: "Kokushimusou"}
				hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &kokushi)
			}
		}
	}
}

func checkKokushiAllTiles(handCount map[string]int) bool {
	kokushi := []string{"H1", "H2", "H3", "H4", "H5", "H6", "H7", "M1", "M9", "P1", "P9", "S1", "S9"}

	//need at least 1 of each tile
	for _, tile := range kokushi {
		if handCount[tile] < 1 {
			return false
		}
	}

	//check have 1 pair of a kokushi tile
	for _, tile := range kokushi {
		if handCount[tile] > 1 {
			return true
		}
	}

	return false
}
func checkKokushiThirteenWait(handParts *types.HandPartBlocks) bool {
	head := getHead(handParts.Menzen)
	if head != nil && handParts.Agari[:2] == head[0][:2] {
		return true
	}
	return false
}

func chuurenpoutou(hand *types.WinningHand) {
	if !checkOpenHand(hand) {

		var suit byte
		tileCount := make(map[byte]int)
		for _, block := range hand.Hand {
			for _, tile := range block {
				if suit != 0 && tile[0] != suit {
					return
				}
				tileCount[tile[1]]++
				suit = tile[0]
			}
		}

	}

}

func daisangen(hand *types.WinningHand) {
	var haku, hatsu, chun bool
	for _, block := range hand.Hand {
		if len(block) >= 3 {
			switch block[0] {
			case "H5":
				haku = true
			case "H6":
				hatsu = true
			case "H7":
				chun = true
			}
		}
	}

	if haku && hatsu && chun {
		daisangen := types.YakumanComponet{Yakuman: 1, Title: "Daisangen"}
		hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &daisangen)
	}
}

func shousuushi(hand *types.WinningHand) {

}

func daisuushi(hand *types.WinningHand) {

}

func tsuuiisou(hand *types.WinningHand) {
	var jihai = map[string]bool{
		"H1": true,
		"H2": true,
		"H3": true,
		"H4": true,
		"H5": true,
		"H6": true,
		"H7": true,
	}

	for _, block := range hand.Hand {
		for _, tile := range block {
			if !jihai[tile] {
				return
			}
		}
	}
	tsuuiisou := types.YakumanComponet{Yakuman: 1, Title: "Tsuuiisou"}
	hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &tsuuiisou)
}

func chinroutou(hand *types.WinningHand) {
	var terminal = map[string]bool{
		"M1": true,
		"M9": true,
		"P1": true,
		"P9": true,
		"S1": true,
		"S9": true,
	}

	for _, block := range hand.Hand {
		for _, tile := range block {
			if !terminal[tile] {
				return
			}
		}
	}
	chinroutou := types.YakumanComponet{Yakuman: 1, Title: "Chinroutou"}
	hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &chinroutou)
}

func ryuuiisou(hand *types.WinningHand) {
	var terminal = map[string]bool{
		"S2": true,
		"S3": true,
		"S4": true,
		"S6": true,
		"S8": true,
		"H6": true,
	}

	for _, block := range hand.Hand {
		for _, tile := range block {
			if !terminal[tile] {
				return
			}
		}
	}
	ryuuiisou := types.YakumanComponet{Yakuman: 1, Title: "Ryuuiisou"}
	hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &ryuuiisou)
}

func suukantsu(hand *types.WinningHand) {
	if len(hand.HandParts.Kan)+len(hand.HandParts.Ankan) >= 4 {
		suukantsu := types.YakumanComponet{Yakuman: 1, Title: "Suukantsu"}
		hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &suukantsu)
	}
}
