package scoring

import (
	"fmt"

	"github.com/KainoaGardner/webMahjongCalc/types"
	"github.com/KainoaGardner/webMahjongCalc/utils"
)

func getYaku(hand *types.WinningHand) {

	//1 han
	riichi(hand)
	menzentsumo(hand)
	ippatsu(hand)
	pinfu(hand)
	iipeikou(hand)
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
	sanankou(hand)
	honroutou(hand)
	sankantsu(hand)

	//3 han
	ryanpeikou(hand)
	honitsu(hand)
	junchan(hand)

	//6 han
	chinitsu(hand)

	//dora
	dora(hand)
	akaDora(hand)
	uraDora(hand)

	hand.HandScore.Han = getTotalHan(hand)
}

func getTotalHan(hand *types.WinningHand) int {
	var totalHan int
	for _, han := range hand.HandScore.YakuList {
		totalHan += han.Han
	}
	return totalHan
}

func riichi(hand *types.WinningHand) {
	if utils.CheckOpenHand(hand) {
		return
	}

	if hand.ScoringParts.Riichi {
		riichi := types.YakuComponet{Han: 1, Title: "Riichi"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &riichi)
	}
}

func menzentsumo(hand *types.WinningHand) {
	if utils.CheckOpenHand(hand) {
		return
	}
	if hand.ScoringParts.Tsumo {
		tsumo := types.YakuComponet{Han: 1, Title: "Menzenchin Tsumohou"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &tsumo)
	}
}

func wriichi(hand *types.WinningHand) {
	if utils.CheckOpenHand(hand) {
		return
	}

	if !hand.ScoringParts.Riichi && hand.ScoringParts.Wriichi {
		wriichi := types.YakuComponet{Han: 2, Title: "Double Riichi"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &wriichi)
	}
}

func ippatsu(hand *types.WinningHand) {
	if utils.CheckOpenHand(hand) {
		return
	}

	if (hand.ScoringParts.Riichi || hand.ScoringParts.Wriichi) && hand.ScoringParts.Ippatsu {
		ippatsu := types.YakuComponet{Han: 1, Title: "Ippatsu"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &ippatsu)
	}
}

func pinfu(hand *types.WinningHand) {
	if utils.CheckOpenHand(hand) || len(hand.HandParts.Ankan) > 0 {
		return
	}

	var ryanmen bool
	for _, block := range hand.HandParts.Menzen {
		if len(block) > 2 {

			if utils.CheckShuntsuBlock(block) {
				if hand.HandParts.Agari == block[0] || hand.HandParts.Agari == block[2] {
					ryanmen = true
				}
			} else {
				return
			}
		}
	}

	if !ryanmen {
		return
	}

	var fuTiles = map[string]bool{
		"H5": true,
		"H6": true,
		"H7": true,
	}

	fuTiles[hand.ScoringParts.Bakaze] = true
	fuTiles[hand.ScoringParts.Jikaze] = true

	head, ok := utils.GetHead(hand.HandParts.Menzen)
	if !ok {
		return
	}

	if fuTiles[head[0]] {
		return
	}

	pinfu := types.YakuComponet{Han: 1, Title: "Pinfu"}
	hand.HandScore.YakuList = append(hand.HandScore.YakuList, &pinfu)

}

func iipeikou(hand *types.WinningHand) {
	if utils.CheckOpenHand(hand) {
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
		if utils.CheckShuntsuBlock(block) {
			shuntsuString := utils.GetBlockString(block)
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
			if utils.CheckOpenHand(hand) {
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
		if utils.CheckKoutsuBlock(block) {
			koutsuString := utils.GetBlockString(block)[:4]
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
		if utils.CheckShuntsuBlock(block) {
			shuntsuString := utils.GetBlockString(block)
			shuntsuMap[shuntsuString] = true
		}
	}

	for _, suit := range []string{"M", "S", "P"} {
		a := fmt.Sprintf("%s123", suit)
		b := fmt.Sprintf("%s456", suit)
		c := fmt.Sprintf("%s789", suit)
		if shuntsuMap[a] && shuntsuMap[b] && shuntsuMap[c] {
			ittsuu := types.YakuComponet{Han: 2, Title: "Ikkitsuukan"}

			if utils.CheckOpenHand(hand) {
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

	if utils.CheckOpenHand(hand) {
		chanta.Han = 1
	}
	hand.HandScore.YakuList = append(hand.HandScore.YakuList, &chanta)

}

func toitoi(hand *types.WinningHand) {
	if len(hand.Hand) != 5 {
		return
	}
	for _, block := range hand.Hand {
		if len(block) != 2 {
			if !utils.CheckKoutsuBlock(block) {
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

func sanankou(hand *types.WinningHand) {
	ankouCount := utils.GetMenzenKoutsuCount(hand.HandParts.Menzen) + len(hand.HandParts.Ankan)

	var winTileInKoutsu bool
	var winTileInShuntsu bool
	for _, block := range hand.HandParts.Menzen {
		if utils.CheckKoutsuBlock(block) {
			if block[0][:2] == hand.HandParts.Agari[:2] {
				winTileInKoutsu = true
			}
		} else {
			for _, tile := range block {
				if tile == hand.HandParts.Agari {
					winTileInShuntsu = true
				}
			}

		}
	}

	for _, block := range hand.HandParts.Ankan {
		if block[0][:2] == hand.HandParts.Agari[:2] {
			winTileInKoutsu = true
		}
	}

	if ankouCount == 3 {
		if hand.ScoringParts.Tsumo || !winTileInKoutsu || (winTileInKoutsu && winTileInShuntsu) {

			sanankou := types.YakuComponet{Han: 2, Title: "Sanankou"}
			hand.HandScore.YakuList = append(hand.HandScore.YakuList, &sanankou)
		}
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
	hand.HandScore.YakuList = utils.RemoveYaku(hand.HandScore.YakuList, "Chanta")
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

	honitsu := types.YakuComponet{Han: 3, Title: "Honitsu"}
	if utils.CheckOpenHand(hand) {
		honitsu.Han = 2
	}
	hand.HandScore.YakuList = append(hand.HandScore.YakuList, &honitsu)
}

func junchan(hand *types.WinningHand) {
	var terminalHonors = map[string]bool{
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
	junchan := types.YakuComponet{Han: 3, Title: "Junchan"}

	if utils.CheckOpenHand(hand) {
		junchan.Han = 2
	}
	hand.HandScore.YakuList = append(hand.HandScore.YakuList, &junchan)
	hand.HandScore.YakuList = utils.RemoveYaku(hand.HandScore.YakuList, "Chanta")

}

func ryanpeikou(hand *types.WinningHand) {
	if utils.CheckOpenHand(hand) {
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

	if peikouCount == 2 {
		ryanpeikou := types.YakuComponet{Han: 3, Title: "Ryanpeikou"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &ryanpeikou)
		hand.HandScore.YakuList = utils.RemoveYaku(hand.HandScore.YakuList, "Iipeikou")
	}

}
func chinitsu(hand *types.WinningHand) {
	for i := 1; i < len(hand.Hand); i++ {
		//it no honor and suits are not equal
		if hand.Hand[i][0][0] == 'H' || hand.Hand[i-1][0][0] == 'H' || hand.Hand[i][0][0] != hand.Hand[i-1][0][0] {
			return
		}
	}

	chinitsu := types.YakuComponet{Han: 6, Title: "Chinitsu"}
	if utils.CheckOpenHand(hand) {
		chinitsu.Han = 5
	}
	hand.HandScore.YakuList = append(hand.HandScore.YakuList, &chinitsu)
	hand.HandScore.YakuList = utils.RemoveYaku(hand.HandScore.YakuList, "Honitsu")
}

func dora(hand *types.WinningHand) {
	dora := make(map[string]bool)
	var doraCount int
	for _, doraTile := range hand.ScoringParts.Dora {
		dora[doraTile[:2]] = true
	}

	for _, block := range hand.Hand {
		for _, tile := range block {
			if dora[tile[:2]] {
				doraCount++
			}
		}
	}

	if doraCount > 0 {
		doraYaku := types.YakuComponet{Han: doraCount, Title: "Dora"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &doraYaku)
	}
}

func uraDora(hand *types.WinningHand) {
	if !hand.ScoringParts.Riichi && !hand.ScoringParts.Wriichi {
		return
	}

	dora := make(map[string]bool)
	var doraCount int
	for _, doraTile := range hand.ScoringParts.Uradora {
		dora[doraTile[:2]] = true
	}

	for _, block := range hand.Hand {
		for _, tile := range block {
			if dora[tile[:2]] {
				doraCount++
			}
		}
	}

	if doraCount > 0 {
		uradoraYaku := types.YakuComponet{Han: doraCount, Title: "Uradora"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &uradoraYaku)
	}

}

func akaDora(hand *types.WinningHand) {
	var doraCount int

	for _, block := range hand.Hand {
		for _, tile := range block {
			if len(tile) == 3 && tile[2] == 'A' {
				doraCount++
			}
		}
	}

	if doraCount > 0 {
		akadoraYaku := types.YakuComponet{Han: doraCount, Title: "Akadora"}
		hand.HandScore.YakuList = append(hand.HandScore.YakuList, &akadoraYaku)
	}
}
