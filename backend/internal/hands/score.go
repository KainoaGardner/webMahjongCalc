package hands

import (
	"fmt"
	"math"

	"github.com/KainoaGardner/webMahjongCalc/types"
	"github.com/KainoaGardner/webMahjongCalc/utils"
)

func GetMaxScoreWin(validHands [][][]string, hand *types.PostHandScore) (*types.WinningHand, error) {
	var winningHand *types.WinningHand

	for _, validHand := range validHands {
		var currentHand types.WinningHand
		currentHand.Hand = validHand

		currentHand.Open = checkOpenHand(hand.Hand)
		currentHand.ScoringParts = hand.ScoringParts

		var score types.Score
		currentHand.HandScore = &score

		err := getHandScore(&currentHand)
		if err != nil {
			return nil, err
		}
		fmt.Println(currentHand)
	}
	return winningHand, nil
}

func checkOpenHand(hand *types.HandParts) bool {
	if len(hand.Chi) > 0 || len(hand.Pon) > 0 || len(hand.Kan) > 0 {
		return true
	}

	return false
}

func getHandScore(currentHand *types.WinningHand) error {
	// han, yakuList := GetYaku(currentHand)
	// fu, fuList := GetFu(currentHand)
	//yakuman
	han := 4
	fu := 40
	score, scoreType, oyaPayment, koPayment, err := getScore(han, fu, currentHand.ScoringParts)
	if err != nil {
		return err
	}

	currentHand.HandScore.Han = han
	currentHand.HandScore.Fu = fu
	currentHand.HandScore.Score = score
	currentHand.HandScore.ScoreType = scoreType
	currentHand.HandScore.OyaPayment = oyaPayment
	currentHand.HandScore.KoPayment = koPayment
	// currentHand.HandScore.YakuList = yakuList
	// currentHand.HandScore.FuList = fuList

	// fmt.Printf("Score: %d, ScoreType: %s, OyaPay %d, KoPay %d\n", score, scoreType, oyaPayment, koPayment)
	return nil
}

func getScore(han int, fu int, scoringParts *types.HandScoringParts) (int, string, int, int, error) {
	var handScore int
	var basicPoints int
	var scoreType string
	var oyaPayment, koPayment int

	if han == 5 ||
		(han == 4 && fu >= 40) || (han == 3 && fu >= 70) ||
		scoringParts.Kiriage && (han == 4 && fu >= 30) || (han == 3 && fu >= 60) {
		basicPoints = 2000
		scoreType = "mangan"
	} else if han == 6 || han == 7 {
		basicPoints = 3000
		scoreType = "haneman"
	} else if han == 8 || han == 9 || han == 10 {
		basicPoints = 4000
		scoreType = "baiman"
	} else if han == 11 || han == 12 {
		basicPoints = 6000
		scoreType = "sanbaiman"
	} else if han >= 13 {
		basicPoints = 8000
		scoreType = "kazoeYakuman"
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
