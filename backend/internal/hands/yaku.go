package hands

import (
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func getYaku(hand *types.WinningHand) {
	riichi(hand)
	wriichi(hand)
	ippatsu(hand)
	// pinfu(hand)
	iiryanpeikou(hand)
	tanyao(hand)
	yakuhai(hand)

}

func riichi(hand *types.WinningHand) {
	if checkOpenHand(hand) {
		return
	}

	if hand.ScoringParts.Riichi {
		riichi := types.YakuComponet{Han: 1, Title: "Riichi"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &riichi)
	}
}

func wriichi(hand *types.WinningHand) {
	if checkOpenHand(hand) {
		return
	}

	if !hand.ScoringParts.Riichi && hand.ScoringParts.Wriichi {
		wriichi := types.YakuComponet{Han: 2, Title: "Double Riichi"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &wriichi)
	}
}

func ippatsu(hand *types.WinningHand) {
	if checkOpenHand(hand) {
		return
	}

	if (hand.ScoringParts.Riichi || hand.ScoringParts.Wriichi) && hand.ScoringParts.Ippatsu {
		ippatsu := types.YakuComponet{Han: 1, Title: "Ippatsu"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &ippatsu)
	}
}

func pinfu(hand *types.WinningHand) {

}

func iiryanpeikou(hand *types.WinningHand) {
	if checkOpenHand(hand) {
		return
	}

	shuntsuCount := make(map[string]int)

	for _, block := range hand.HandParts.Menzen {
		if block[0] != block[1] {
			shuntsu := ""
			for _, tile := range block {
				shuntsu += tile[:2]
			}
			shuntsuCount[shuntsu]++
		}
	}

	peikouCount := 0
	for _, amount := range shuntsuCount {
		if amount >= 2 {
			peikouCount++
		}

	}

	if peikouCount == 1 {
		iipeikou := types.YakuComponet{Han: 1, Title: "Iipeikou"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &iipeikou)
	} else if peikouCount == 2 {
		ryanpeikou := types.YakuComponet{Han: 3, Title: "Ryanpeikou"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &ryanpeikou)
	}

}

func tanyao(hand *types.WinningHand) {
	var tanyaohai = map[string]bool{
		"H1": true,
		"H2": true,
		"H3": true,
		"H4": true,
		"H5": true,
		"H6": true,
		"H7": true,
		"M1": true,
		"M9": true,
		"S1": true,
		"S9": true,
		"P1": true,
		"P9": true,
	}

	for _, block := range hand.Hand {
		for _, tile := range block {
			if tanyaohai[tile] {
				return
			}
		}
	}
	tanyao := types.YakuComponet{Han: 1, Title: "Tanyao"}
	hand.HandScore.YakuList = append(hand.HandScore.YakuList, &tanyao)

}

func yakuhai(hand *types.WinningHand) {
	for _, block := range hand.Hand {
		if len(block) >= 3 {
			if block[0] == "H5" {
				haku := types.YakuComponet{Han: 1, Title: "Haku"}
				hand.HandScore.YakuList = append(hand.HandScore.YakuList, &haku)
			}
			if block[0] == "H6" {
				hatsu := types.YakuComponet{Han: 1, Title: "Hatsu"}
				hand.HandScore.YakuList = append(hand.HandScore.YakuList, &hatsu)
			}
			if block[0] == "H7" {
				chun := types.YakuComponet{Han: 1, Title: "Chun"}
				hand.HandScore.YakuList = append(hand.HandScore.YakuList, &chun)
			}
			if block[0] == hand.ScoringParts.Bakaze {
				bakaze := types.YakuComponet{Han: 1, Title: "Bakaze"}
				hand.HandScore.YakuList = append(hand.HandScore.YakuList, &bakaze)
			}
			if block[0] == hand.ScoringParts.Jikaze {
				jikaze := types.YakuComponet{Han: 1, Title: "Jikaze"}
				hand.HandScore.YakuList = append(hand.HandScore.YakuList, &jikaze)
			}
		}
	}
}
