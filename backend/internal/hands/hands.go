package hands

import (
	"github.com/KainoaGardner/webMahjongCalc/internal/scoring"
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func GetHandScore(hand *types.PostHandScore) (*types.ReturnHandScore, error) {
	var scoreResult types.ReturnHandScore
	err := checkValidData(hand)
	if err != nil {
		return nil, err
	}

	validHands, err := GetValidHands(hand.Hand)
	if err != nil {
		return nil, err
	}

	winningHand, err := scoring.GetMaxScoreWin(validHands, hand)
	if err != nil {
		return nil, err
	}

	scoreResult.Hand = winningHand.HandParts
	scoreResult.HandScore = winningHand.HandScore

	return &scoreResult, nil

}
