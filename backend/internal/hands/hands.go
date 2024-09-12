package hands

import (
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func GetHandScore(hand *types.PostHandScore) (*types.ReturnHandScore, error) {
	var scoreResult types.ReturnHandScore
	err := CheckValidData(hand)
	if err != nil {
		return nil, err
	}

	winningHands, err := GetWinningHands(hand)
	if err != nil {
		return nil, err
	}

	winningHand, handScore, err := getMaxScoreWin(winningHands)
	if err != nil {
		return nil, err
	}
	scoreResult.Hand = winningHand
	scoreResult.HandScore = handScore

	return &scoreResult, nil

}

func getMaxScoreWin(winningHands *[]types.HandPartsBlocks) (*types.HandPartsBlocks, *types.Score, error) {
	var winningHand types.HandPartsBlocks
	var handScore types.Score

	return &winningHand, &handScore, nil
}
