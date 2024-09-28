package hands

import (
	"fmt"
	"math"

	"github.com/KainoaGardner/webMahjongCalc/types"
	"github.com/KainoaGardner/webMahjongCalc/utils"
)

func getMaxScoreWin(validHands []types.WinningHand, hand *types.PostHandScore) (*types.WinningHand, error) {
	var winningHand *types.WinningHand

	for _, validHand := range validHands {
		// currentHand.Hand = validHand

		validHand.ScoringParts = hand.ScoringParts

		var score types.Score
		validHand.HandScore = &score

		err := getHandScore(&validHand)
		if err != nil {
			return nil, err
		}
	}
	return winningHand, nil
}

func getHandScore(currentHand *types.WinningHand) error {
	//yakuman
	getYakuman(currentHand)

	getYaku(currentHand)
	for _, yaku := range currentHand.HandScore.YakuList {
		fmt.Println(yaku)

	}

	// if currentHand.Yakuman != 0 || len(currentHand.HandScore.YakumanList) != 0 {
	//
	// }else

	// han, yakuList := GetYaku(currentHand)
	// GetFu(currentHand)

	// fu := 3
	// han := 4
	// yakuman := 3

	// score, scoreType, oyaPayment, koPayment, err := getScore(currentHand.HandScore.Han, currentHand.HandScore.Fu, currentHand.HandScore.Yakuman, currentHand.ScoringParts)
	// if err != nil {
	// 	return err
	// }
	//
	// // currentHand.HandScore.Fu = fu
	// // currentHand.HandScore.FuList = fuList
	//
	// currentHand.HandScore.Score = score
	// currentHand.HandScore.ScoreType = scoreType
	// currentHand.HandScore.OyaPayment = oyaPayment
	// currentHand.HandScore.KoPayment = koPayment
	// // currentHand.HandScore.YakuList = yakuList

	// fmt.Printf("Score: %d, ScoreType: %s, OyaPay %d, KoPay %d\n", score, scoreType, oyaPayment, koPayment)
	return nil
}

// update for yakuman!!!!!!
func getScore(han int, fu int, yakuman int, scoringParts *types.HandScoringParts) (int, string, int, int, error) {
	var handScore int
	var basicPoints int
	var scoreType string
	var oyaPayment, koPayment int

	if han == 5 ||
		(han == 4 && fu >= 40) || (han == 3 && fu >= 70) ||
		scoringParts.Kiriage && (han == 4 && fu >= 30) || (han == 3 && fu >= 60) {
		basicPoints = 2000
		scoreType = "Mangan"
	} else if han == 6 || han == 7 {
		basicPoints = 3000
		scoreType = "Haneman"
	} else if han == 8 || han == 9 || han == 10 {
		basicPoints = 4000
		scoreType = "Baiman"
	} else if han == 11 || han == 12 {
		basicPoints = 6000
		scoreType = "Sanbaiman"
	} else if han >= 13 {
		basicPoints = 8000
		scoreType = "Kazoe Yakuman"
	} else {
		basicPoints = fu * int(math.Pow(2, float64(2+han)))
	}

	if scoringParts.Tsumo {
		if scoringParts.Oya {
			koPayment = utils.CeilHundred(basicPoints * 2)
			koPayment += scoringParts.Honba * 100
			handScore = koPayment * 3
		} else {
			koPayment = utils.CeilHundred(basicPoints)
			oyaPayment = utils.CeilHundred(basicPoints * 2)

			koPayment += scoringParts.Honba * 100
			oyaPayment += scoringParts.Honba * 100

			handScore = oyaPayment + koPayment*2
		}
	} else if scoringParts.Ron {
		if scoringParts.Oya {
			handScore = utils.CeilHundred(basicPoints * 6)
			handScore += scoringParts.Honba * 300
		} else {
			handScore = utils.CeilHundred(basicPoints * 4)
			handScore += scoringParts.Honba * 300
		}
	} else {
		return 0, "", 0, 0, fmt.Errorf("Must have agari type (ron or tsumo)")
	}

	return handScore, scoreType, oyaPayment, koPayment, nil
}
