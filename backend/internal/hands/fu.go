package hands

import (
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func GetFu(currentHand *types.WinningHand) {
	FuAdded := types.FuComponet{Fu: 20, Title: "Fuutei"}
	currentHand.HandScore.FuList = append(currentHand.HandScore.FuList, &FuAdded)

	menzenKafu(currentHand)

}

func menzenKafu(currentHand *types.WinningHand) {
	if !checkOpenHand(currentHand) && currentHand.ScoringParts.Ron {
		FuAdded := types.FuComponet{Fu: 10, Title: "Menzen Kafu"}
		currentHand.HandScore.FuList = append(currentHand.HandScore.FuList, &FuAdded)
	}
}
