package scoring

import (
	"fmt"
	"math"

	"github.com/KainoaGardner/webMahjongCalc/types"
	"github.com/KainoaGardner/webMahjongCalc/utils"
)

func GetMaxScoreWin(validHands []types.WinningHand, hand *types.PostHandScore) (*types.WinningHand, error) {
	var winningHand *types.WinningHand

	for _, validHand := range validHands {
		validHand.ScoringParts = hand.ScoringParts

		err := getHandScore(&validHand)
		if err != nil {
			return nil, err
		}
	}

	for _, validHand := range validHands {
		if winningHand == nil || validHand.HandScore.Score > winningHand.HandScore.Score {
			winningHand = &validHand
		}
	}
	return winningHand, nil
}

func getHandScore(currentHand *types.WinningHand) error {
	//yakuman
	getYakuman(currentHand)
	if currentHand.HandScore.Yakuman == 0 {
		getYaku(currentHand)
		if currentHand.HandScore.Han < 5 {
			getFu(currentHand)
		}
	}

	err := getScore(currentHand)
	if err != nil {
		return err
	}

	return nil
}

// update for yakuman!!!!!!
func getScore(hand *types.WinningHand) error {
	var handScore int
	var basicPoints int
	var scoreType string
	var oyaPayment, koPayment int

	han := hand.HandScore.Han
	fu := hand.HandScore.Fu
	yakuman := hand.HandScore.Yakuman

	if checkNoYaku(hand) && yakuman == 0 {
		basicPoints = 0
		scoreType = "No Yaku"
	} else if han == 5 ||
		(han == 4 && fu >= 40) || (han == 3 && fu >= 70) ||
		hand.ScoringParts.Kiriage && (han == 4 && fu >= 30) || (han == 3 && fu >= 60) {
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
	} else if yakuman > 0 {
		basicPoints = 8000 * yakuman
		switch yakuman {
		case 2:
			scoreType = fmt.Sprintf("Double Yakuman")
		case 3:
			scoreType = fmt.Sprintf("Triple Yakuman")
		case 4:
			scoreType = fmt.Sprintf("Quadruple Yakuman")
		case 5:
			scoreType = fmt.Sprintf("Quintuple Yakuman")
		case 6:
			scoreType = fmt.Sprintf("Sextuple Yakuman")
		default:
			scoreType = fmt.Sprintf("Yakuman")
		}
	} else {
		basicPoints = fu * int(math.Pow(2, float64(2+han)))
	}

	if hand.ScoringParts.Tsumo {
		if hand.ScoringParts.Oya {
			koPayment = utils.CeilHundred(basicPoints * 2)
			koPayment += hand.ScoringParts.Honba * 100
			handScore = koPayment * 3
		} else {
			koPayment = utils.CeilHundred(basicPoints)
			oyaPayment = utils.CeilHundred(basicPoints * 2)

			koPayment += hand.ScoringParts.Honba * 100
			oyaPayment += hand.ScoringParts.Honba * 100

			handScore = oyaPayment + koPayment*2
		}
	} else if hand.ScoringParts.Ron {
		if hand.ScoringParts.Oya {
			handScore = utils.CeilHundred(basicPoints * 6)
			handScore += hand.ScoringParts.Honba * 300
		} else {
			handScore = utils.CeilHundred(basicPoints * 4)
			handScore += hand.ScoringParts.Honba * 300
		}
	} else {
		return fmt.Errorf("Must have agari type (ron or tsumo)")
	}

	handScore += hand.ScoringParts.RiichiBou * 1000

	hand.HandScore.Score = handScore
	hand.HandScore.ScoreType = scoreType
	hand.HandScore.OyaPayment = oyaPayment
	hand.HandScore.KoPayment = koPayment

	return nil
}

func checkNoYaku(hand *types.WinningHand) bool {
	for _, yaku := range hand.HandScore.YakuList {
		if yaku.Title != "Dora" && yaku.Title != "Akadora" && yaku.Title != "Uradora" {
			return false
		}
	}
	return true
}
