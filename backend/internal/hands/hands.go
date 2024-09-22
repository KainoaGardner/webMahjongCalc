package hands

import (
	"fmt"

	"github.com/KainoaGardner/webMahjongCalc/types"
)

func GetHandScore(hand *types.PostHandScore) (*types.ReturnHandScore, error) {
	var scoreResult types.ReturnHandScore
	err := CheckValidData(hand)
	if err != nil {
		return nil, err
	}

	validHands, err := GetValidHands(hand.Hand)
	if err != nil {
		return nil, err
	}

	winningHand, err := GetMaxScoreWin(validHands, hand)
	if err != nil {
		return nil, err
	}

	fmt.Println(winningHand)
	// scoreResult.Hand = winningHand
	// scoreResult.HandScore = handScore
	//
	return &scoreResult, nil

}
