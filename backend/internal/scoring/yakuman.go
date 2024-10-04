package scoring

import (
	"github.com/KainoaGardner/webMahjongCalc/types"
	"github.com/KainoaGardner/webMahjongCalc/utils"
)

func getYakuman(hand *types.WinningHand) {
	tenhou(hand)
	chiihou(hand)
	suuankou(hand)
	kokushimusou(hand)
	chuurenpoutou(hand)
	daisangen(hand)
	daishousuushi(hand)
	tsuuiisou(hand)
	chinroutou(hand)
	ryuuiisou(hand)
	suukantsu(hand)

	hand.HandScore.Yakuman = getTotalYakuman(hand)

}
func getTotalYakuman(hand *types.WinningHand) int {
	var totalYakuman int
	for _, yakuman := range hand.HandScore.YakumanList {
		totalYakuman += yakuman.Yakuman
	}
	return totalYakuman
}

func tenhou(hand *types.WinningHand) {
	if utils.CheckOpenHand(hand) {
		return
	}
	if hand.ScoringParts.Tenhou {
		tenhou := types.YakumanComponet{Yakuman: 1, Title: "Tenhou"}
		hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &tenhou)
	}

}

func chiihou(hand *types.WinningHand) {
	if utils.CheckOpenHand(hand) {
		return
	}
	if hand.ScoringParts.Tenhou {
		chiihou := types.YakumanComponet{Yakuman: 1, Title: "Chiihou"}
		hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &chiihou)
	}

}

func suuankou(hand *types.WinningHand) {
	koutsuCount := utils.GetMenzenKoutsuCount(hand.HandParts.Menzen) + len(hand.HandParts.Ankan)
	if koutsuCount >= 4 {

		head, ok := utils.GetHead(hand.HandParts.Menzen)
		if !ok {
			return
		}

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
	if utils.CheckOpenHand(hand) {
		return
	}
	unsortedHand := utils.GetUnsortedHand(hand)
	handCount := utils.GetTileMap(unsortedHand)

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
	head, ok := utils.GetHead(handParts.Menzen)
	if !ok {
		return false
	}
	if head != nil && handParts.Agari[:2] == head[0][:2] {
		return true
	}
	return false
}

func chuurenpoutou(hand *types.WinningHand) {
	if utils.CheckOpenHand(hand) {
		return
	}

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

	for tile, amount := range tileCount {
		if tile == '1' || tile == '9' {
			if amount < 3 {
				return
			} else {
				if amount < 1 {
					return
				}
			}
		}
	}

	chuuren := types.YakumanComponet{Yakuman: 1, Title: "Chuuren Poutou"}
	junseichuuren := types.YakumanComponet{Yakuman: 2, Title: "Junsei Chuuren Poutou"}

	agariNumber := hand.HandParts.Agari[1]
	if agariNumber == '1' || agariNumber == '9' {
		if tileCount[agariNumber] == 4 {
			hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &junseichuuren)
		} else {
			hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &chuuren)
		}

	} else {
		if tileCount[agariNumber] == 2 {
			hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &junseichuuren)
		} else {
			hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &chuuren)
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

func daishousuushi(hand *types.WinningHand) {
	var kazeTiles = map[string]bool{
		"H1": true,
		"H2": true,
		"H3": true,
		"H4": true,
	}

	kaze := [][]string{}
	for _, block := range hand.Hand {
		if !kazeTiles[block[0]] {
			return
		}

		kaze = append(kaze, block)

	}

	if len(kaze) == 4 {
		for _, block := range kaze {
			if len(block) == 2 {
				shousuushi := types.YakumanComponet{Yakuman: 1, Title: "Shousuushi"}
				hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &shousuushi)
				return
			}
		}
		daisuushi := types.YakumanComponet{Yakuman: 2, Title: "Daisuushi"}
		hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &daisuushi)
	}

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
