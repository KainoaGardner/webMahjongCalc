package hands

import (
	"fmt"

	"github.com/KainoaGardner/webMahjongCalc/types"
)

func getYaku(hand *types.WinningHand) {

	//1 han
	riichi(hand)
	menzentsumo(hand)
	ippatsu(hand)
	// pinfu(hand)
	iiryanpeikou(hand) //check both iipeikou and ryanpeikou
	tanyao(hand)
	yakuhai(hand)
	haitei(hand)
	houtei(hand)
	rinshan(hand)
	chankan(hand)

	//2han
	wriichi(hand)
	chiitoitsu(hand)
	sanshokudoujun(hand)
	sanshokudoukou(hand)
	ittsuu(hand)
	chanta(hand)
	toitoi(hand)
	shousangen(hand)
	// sanankou(hand) //NOT DONE
	honroutou(hand)
	sankantsu(hand)

	//3 han
	// ryanpeikou(hand)
	honitsu(hand)
	// junchan(hand)

	//6 han
	// chinittsu(hand)

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

func menzentsumo(hand *types.WinningHand) {
	if checkOpenHand(hand) {
		return
	}
	if hand.ScoringParts.Tsumo {
		tsumo := types.YakuComponet{Han: 1, Title: "Menzenchin Tsumohou"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &tsumo)
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

func haitei(hand *types.WinningHand) {
	if hand.ScoringParts.Haitei && hand.ScoringParts.Tsumo {
		haitei := types.YakuComponet{Han: 1, Title: "Haitei Raoyue"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &haitei)
	}

}

func houtei(hand *types.WinningHand) {
	if hand.ScoringParts.Houtei && hand.ScoringParts.Ron {
		houtei := types.YakuComponet{Han: 1, Title: "Houtei Raoyui"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &houtei)
	}

}

func rinshan(hand *types.WinningHand) {
	if len(hand.HandParts.Kan)+len(hand.HandParts.Ankan) == 0 {
		return
	}

	if hand.ScoringParts.Rinshan {
		rinshan := types.YakuComponet{Han: 1, Title: "Rinshan Kaihou"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &rinshan)
	}

}

func chankan(hand *types.WinningHand) {
	if hand.ScoringParts.Ron && hand.ScoringParts.Chankan {
		chankan := types.YakuComponet{Han: 1, Title: "Chankan"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &chankan)
	}
}

func chiitoitsu(hand *types.WinningHand) {
	if len(hand.Hand) == 7 {
		chiitoitsu := types.YakuComponet{Han: 2, Title: "Chiitoitsu"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &chiitoitsu)
	}

}

func sanshokudoujun(hand *types.WinningHand) {
	shuntsuMap := make(map[string]map[byte]bool)

	for _, block := range hand.Hand {
		if checkShuntsuBlock(block) {
			shuntsuString := getBlockString(block)
			if shuntsuMap[shuntsuString[1:]] == nil {
				var suitMap = map[byte]bool{'M': false, 'P': false, 'S': false}
				shuntsuMap[shuntsuString[1:]] = suitMap
			}
			shuntsuMap[shuntsuString[1:]][shuntsuString[0]] = true
		}
	}

	for _, suits := range shuntsuMap {
		if suits['M'] && suits['S'] && suits['P'] {
			sanshoku := types.YakuComponet{Han: 2, Title: "Sanshoku Doujun"}
			//open 1 han
			if checkOpenHand(hand) {
				sanshoku.Han = 1
			}
			hand.HandScore.YakuList = append(hand.HandScore.YakuList, &sanshoku)
			return
		}

	}

}

func sanshokudoukou(hand *types.WinningHand) {
	koutsuMap := make(map[string]map[byte]bool)

	for _, block := range hand.Hand {
		if checkKoutsuBlock(block) {
			koutsuString := getBlockString(block)[:4]
			if koutsuMap[koutsuString[1:]] == nil {
				var suitMap = map[byte]bool{'M': false, 'P': false, 'S': false}
				koutsuMap[koutsuString[1:]] = suitMap
			}
			koutsuMap[koutsuString[1:]][koutsuString[0]] = true
		}
	}

	for _, suits := range koutsuMap {
		if suits['M'] && suits['S'] && suits['P'] {
			sanshoku := types.YakuComponet{Han: 2, Title: "Sanshoku Doukou"}
			hand.HandScore.YakuList = append(hand.HandScore.YakuList, &sanshoku)
			return
		}

	}

}

func ittsuu(hand *types.WinningHand) {
	shuntsuMap := make(map[string]bool)

	for _, block := range hand.Hand {
		if checkShuntsuBlock(block) {
			shuntsuString := getBlockString(block)
			shuntsuMap[shuntsuString] = true
		}
	}

	for _, suit := range []string{"M", "S", "P"} {
		a := fmt.Sprintf("%s123", suit)
		b := fmt.Sprintf("%s456", suit)
		c := fmt.Sprintf("%s789", suit)
		if shuntsuMap[a] && shuntsuMap[b] && shuntsuMap[c] {
			ittsuu := types.YakuComponet{Han: 2, Title: "Ikkitsuukan"}

			if checkOpenHand(hand) {
				ittsuu.Han = 1
			}
			hand.HandScore.YakuList = append(hand.HandScore.YakuList, &ittsuu)
			return
		}
	}

}

func chanta(hand *types.WinningHand) {
	var terminalHonors = map[string]bool{
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
		var terminalFound bool
		for _, tile := range block {
			if terminalHonors[tile] {
				terminalFound = true
			}
		}
		if !terminalFound {
			return
		}
	}
	chanta := types.YakuComponet{Han: 2, Title: "Chanta"}

	if checkOpenHand(hand) {
		chanta.Han = 1
	}
	hand.HandScore.YakuList = append(hand.HandScore.YakuList, &chanta)

}

func toitoi(hand *types.WinningHand) {
	for _, block := range hand.Hand {
		if len(block) != 2 {
			if !checkKoutsuBlock(block) {
				return
			}
		}
	}
	toitoi := types.YakuComponet{Han: 2, Title: "Toitoihou"}
	hand.HandScore.YakuList = append(hand.HandScore.YakuList, &toitoi)

}

func shousangen(hand *types.WinningHand) {
	if len(hand.Hand) != 5 {
		return
	}

	var haku, hatsu, chun bool

	for _, block := range hand.Hand {
		switch block[0] {
		case "H5":
			haku = true
		case "H6":
			hatsu = true
		case "H7":
			chun = true
		}

	}

	if haku && hatsu && chun {
		shousangen := types.YakuComponet{Han: 2, Title: "Shousangen"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &shousangen)
	}
}

// not done HARD have to know wait i think
func sanankou(hand *types.WinningHand) {
	ankouCount := getMenzenKoutsuCount(hand.HandParts.Menzen) + len(hand.HandParts.Ankan)

	if ankouCount == 3 {
		sanankou := types.YakuComponet{Han: 2, Title: "Sanankou"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &sanankou)
	}

}

func honroutou(hand *types.WinningHand) {
	var terminalHonors = map[string]bool{
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
			if !terminalHonors[tile] {
				return
			}
		}
	}
	honroutou := types.YakuComponet{Han: 2, Title: "Honroutou"}
	hand.HandScore.YakuList = append(hand.HandScore.YakuList, &honroutou)
	hand.HandScore.YakuList = removeYaku(hand.HandScore.YakuList, "Chanta")
}

func sankantsu(hand *types.WinningHand) {
	kanCount := len(hand.HandParts.Kan) + len(hand.HandParts.Ankan)

	if kanCount == 3 {
		sankantsu := types.YakuComponet{Han: 2, Title: "Sankantsu"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &sankantsu)
	}

}

func honitsu(hand *types.WinningHand) {
	for i := 1; i < len(hand.Hand); i++ {
		//it no honor and suits are not equal
		if hand.Hand[i][0][0] != 'H' && hand.Hand[i-1][0][0] != 'H' && hand.Hand[i][0][0] != hand.Hand[i-1][0][0] {
			return
		}
	}

	honitsu := types.YakuComponet{Han: 3, Title: "Honittsu"}
	if checkOpenHand(hand) {
		honitsu.Han = 2
	}
	hand.HandScore.YakuList = append(hand.HandScore.YakuList, &honitsu)
}
