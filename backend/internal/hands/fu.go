package hands

import (
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func getFu(currentHand *types.WinningHand) {

	//chiitoitsu
	if len(currentHand.Hand) == 7 {
		FuAdded := types.FuComponet{Fu: 25, Title: "Chiitoitsu"}
		currentHand.HandScore.FuList = append(currentHand.HandScore.FuList, &FuAdded)

	} else {
		FuAdded := types.FuComponet{Fu: 20, Title: "Fuutei"}
		currentHand.HandScore.FuList = append(currentHand.HandScore.FuList, &FuAdded)

		// pairFu(currentHand)
		// koutsuKantsuFu(currentHand)
		// machiFu(currentHand)
		//
		// menzenKafu(currentHand)
		// tsumoFu(currentHand)
		//
		// openPinfu(currentHand)
	}
}

func menzenKafu(currentHand *types.WinningHand) {
	if !checkOpenHand(currentHand) && currentHand.ScoringParts.Ron {
		FuAdded := types.FuComponet{Fu: 10, Title: "Menzen Kafu"}
		currentHand.HandScore.FuList = append(currentHand.HandScore.FuList, &FuAdded)
	}
}
