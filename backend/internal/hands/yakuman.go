package hands

import (
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func getYakuman(hand *types.WinningHand) error {
	tenhou(hand)
	chiihou(hand)

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

func getMenzenKoutsuCount(menzen [][]string) int {
	var count int
	for _, block := range menzen {
		if checkValidKoutsu(block) {
			count++
		}
	}

	return count
}

func checkValidKoutsu(block []string) bool {
	if len(block) < 3 {
		return false
	}
	for i := 1; i < len(block); i++ {
		if block[i][:2] != block[i-1][:2] {
			return false
		}
	}

	return true
}

func getHead(menzen [][]string) []string {
	for _, block := range menzen {
		if len(block) == 2 {
			return block
		}
	}
	return nil
}

func kokushimusou(hand *types.WinningHand) {
	if !checkOpenHand(hand) {

		if checkKokushiAllTiles(hand.Hand) {
			if checkKokushiThirteenWait(hand.Hand) {
				kokushiJuusanmen := types.YakumanComponet{Yakuman: 2, Title: "Kokushimusou Juusanmenmachi"}
				hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &kokushiJuusanmen)
			} else {
				kokushi := types.YakumanComponet{Yakuman: 1, Title: "Kokushimusou"}
				hand.HandScore.YakumanList = append(hand.HandScore.YakumanList, &kokushi)
			}
		}

	}

}

func checkKokushiAllTiles(hand [][]string) bool {

	return true
}
func checkKokushiThirteenWait(hand [][]string) bool {

	return true
}
