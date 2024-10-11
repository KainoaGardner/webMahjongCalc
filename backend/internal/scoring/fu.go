package scoring

import (
	"fmt"

	"github.com/KainoaGardner/webMahjongCalc/types"
	"github.com/KainoaGardner/webMahjongCalc/utils"
)

func getFu(currentHand *types.WinningHand) {

	//chiitoitsu
	if len(currentHand.Hand) == 7 {
		FuAdded := types.FuComponet{Fu: 25, Title: "Chiitoitsu"}
		currentHand.HandScore.FuList = append(currentHand.HandScore.FuList, &FuAdded)

	} else {
		FuAdded := types.FuComponet{Fu: 20, Title: "Fuutei"}
		currentHand.HandScore.FuList = append(currentHand.HandScore.FuList, &FuAdded)

		yakuhaiHead(currentHand)
		koutsuFu(currentHand)
		kantsuFu(currentHand)

		machiFu(currentHand)

		menzenKafu(currentHand)
		tsumoFu(currentHand)

		openPinfu(currentHand)

		currentHand.HandScore.Fu = utils.CeilTen(getTotalFu(currentHand))
	}
}

func getTotalFu(hand *types.WinningHand) int {
	var totalFu int
	for _, fu := range hand.HandScore.FuList {
		totalFu += fu.Fu
	}
	return totalFu
}

func menzenKafu(hand *types.WinningHand) {
	if !utils.CheckOpenHand(hand) && hand.ScoringParts.Ron {
		FuAdded := types.FuComponet{Fu: 10, Title: "Menzen Kafu"}
		hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)
	}
}

func tsumoFu(hand *types.WinningHand) {
	if utils.CheckYaku(hand.HandScore.YakuList, "Pinfu") {
		return
	}

	if hand.ScoringParts.Tsumo {
		FuAdded := types.FuComponet{Fu: 2, Title: "Tsumo"}
		hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)
	}

}

func yakuhaiHead(hand *types.WinningHand) {
	head, ok := utils.GetHead(hand.Hand)
	if !ok {
		return
	}
	var yakuhaiTiles = map[string]bool{
		"H5": true,
		"H6": true,
		"H7": true,
	}

	yakuhaiTiles[hand.ScoringParts.Bakaze] = true
	yakuhaiTiles[hand.ScoringParts.Jikaze] = true

	if yakuhaiTiles[head[0]] {
		FuAdded := types.FuComponet{Fu: 10, Title: "Yakuhai Pair"}
		hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)
	}

}

func koutsuFu(hand *types.WinningHand) {
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

	for _, block := range hand.HandParts.Menzen {
		if utils.CheckKoutsuBlock(block) {
			//ankou termhonor 8
			if terminalHonors[block[0][:2]] {
				FuAdded := types.FuComponet{Fu: 8, Title: fmt.Sprintf("Ankou")}
				hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)

			} else //ankou simple 4
			{
				FuAdded := types.FuComponet{Fu: 4, Title: fmt.Sprintf("Ankou")}
				hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)

			}
		}

	}
	//all minkou
	for _, block := range hand.HandParts.Pon {
		//minkou terminal
		if terminalHonors[block[0][:2]] {
			FuAdded := types.FuComponet{Fu: 4, Title: fmt.Sprintf("Minkou")}
			hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)

		} else //minkou simple 2
		{
			FuAdded := types.FuComponet{Fu: 2, Title: fmt.Sprintf("Minkou")}
			hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)

		}
	}

}

func kantsuFu(hand *types.WinningHand) {
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

	for _, block := range hand.HandParts.Ankan {
		//ankan terminal 32
		if terminalHonors[block[0][:2]] {
			FuAdded := types.FuComponet{Fu: 32, Title: fmt.Sprintf("Ankan")}
			hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)

		} else //ankan simple
		{
			FuAdded := types.FuComponet{Fu: 16, Title: fmt.Sprintf("Ankan")}
			hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)

		}

	}
	//all minkan
	for _, block := range hand.HandParts.Kan {
		//minkou terminal
		if terminalHonors[block[0][:2]] {
			FuAdded := types.FuComponet{Fu: 16, Title: fmt.Sprintf("Minkan")}
			hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)

		} else //minkou simple 2
		{
			FuAdded := types.FuComponet{Fu: 8, Title: fmt.Sprintf("Minkan")}
			hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)

		}
	}
}

func machiFu(hand *types.WinningHand) {
	// tankiMachi
	if utils.CheckYaku(hand.HandScore.YakuList, "Pinfu") {
		return
	}

	if checkTanki(hand) {
		FuAdded := types.FuComponet{Fu: 2, Title: "Tanki"}
		hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)
	} else if checkKanchan(hand) {
		FuAdded := types.FuComponet{Fu: 2, Title: "Kanchan"}
		hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)
	} else if checkPenchan(hand) {
		FuAdded := types.FuComponet{Fu: 2, Title: "Penchan"}
		hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)

	}

}

func checkTanki(hand *types.WinningHand) bool {
	head, ok := utils.GetHead(hand.Hand)
	if !ok {
		return false
	}
	return hand.HandParts.Agari[:2] == head[0][:2]
}

func checkKanchan(hand *types.WinningHand) bool {
	for _, block := range hand.HandParts.Menzen {
		if utils.CheckShuntsuBlock(block) && block[1] == hand.HandParts.Agari {
			return true
		}
	}

	return false
}

func checkPenchan(hand *types.WinningHand) bool {
	for _, block := range hand.HandParts.Menzen {
		if utils.CheckShuntsuBlock(block) {
			if (block[0][:2] == "M1" || block[0][:2] == "S1" || block[0][:2] == "P1") && block[2] == hand.HandParts.Agari {
				return true
			}
			if (block[2][:2] == "M9" || block[2][:2] == "S9" || block[2][:2] == "P9") && block[0] == hand.HandParts.Agari {
				return true
			}
		}
	}

	return false
}

func openPinfu(hand *types.WinningHand) {
	totalFu := getTotalFu(hand)
	if totalFu == 20 {
		FuAdded := types.FuComponet{Fu: 10, Title: "Open Pinfu"}
		hand.HandScore.FuList = append(hand.HandScore.FuList, &FuAdded)
	}

}
