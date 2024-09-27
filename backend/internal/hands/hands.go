package hands

import (
	"fmt"

	"github.com/KainoaGardner/webMahjongCalc/types"
)

func GetHandScore(hand *types.PostHandScore) (*types.ReturnHandScore, error) {
	var scoreResult types.ReturnHandScore
	err := checkValidData(hand)
	if err != nil {
		return nil, err
	}

	validHands, err := getValidHands(hand.Hand)
	if err != nil {
		return nil, err
	}

	for _, validHand := range validHands {
		fmt.Println(validHand.HandParts)
	}

	winningHand, err := getMaxScoreWin(validHands, hand)
	if err != nil {
		return nil, err
	}

	fmt.Println("last", winningHand)
	// scoreResult.Hand = winningHand
	// scoreResult.HandScore = handScore
	//
	return &scoreResult, nil

}
