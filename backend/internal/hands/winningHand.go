package hands

import (
	"fmt"

	"github.com/KainoaGardner/webMahjongCalc/types"
)

func GetWinningHands(hand *types.PostHandScore) (*[]types.HandPartsBlocks, error) {
	var winningHands []types.HandPartsBlocks
	var handPartsBlocks types.HandPartsBlocks

	possibleHeads, err := getAllHeads(hand)

	sorted, err := sortAllSuit(hand)
	if err != nil {
		return nil, err
	}

	handPartsBlocks.Menzen.Manzu = append(handPartsBlocks.Menzen.Manzu, sorted.Menzen.Manzu)
	handPartsBlocks.Menzen.Souzu = append(handPartsBlocks.Menzen.Souzu, sorted.Menzen.Souzu)
	handPartsBlocks.Menzen.Pinzu = append(handPartsBlocks.Menzen.Pinzu, sorted.Menzen.Pinzu)
	handPartsBlocks.Menzen.Jihai = append(handPartsBlocks.Menzen.Jihai, sorted.Menzen.Jihai)
	winningHands = append(winningHands, handPartsBlocks)

	return &winningHands, nil
}

func sortBySuit(handPart []string, sortedHandPart *types.Suits) error {
	for _, tile := range handPart {
		switch tile[0] {
		case 'H':
			sortedHandPart.Jihai = append(sortedHandPart.Jihai, tile)
		case 'M':
			sortedHandPart.Manzu = append(sortedHandPart.Manzu, tile)
		case 'S':
			sortedHandPart.Souzu = append(sortedHandPart.Souzu, tile)
		case 'P':
			sortedHandPart.Pinzu = append(sortedHandPart.Pinzu, tile)
		default:
			return fmt.Errorf("%s not valid suit", tile)
		}
	}

	return nil
}

func sortAllSuit(hand *types.PostHandScore) (*types.HandSuits, error) {
	var sortedHand types.HandSuits
	err := sortBySuit(hand.Hand.Menzen, &sortedHand.Menzen)
	if err != nil {
		return nil, err
	}
	err = sortBySuit(hand.Hand.Chi, &sortedHand.Chi)
	if err != nil {
		return nil, err
	}
	err = sortBySuit(hand.Hand.Pon, &sortedHand.Pon)
	if err != nil {
		return nil, err
	}
	err = sortBySuit(hand.Hand.Kan, &sortedHand.Kan)
	if err != nil {
		return nil, err
	}
	err = sortBySuit(hand.Hand.Ankan, &sortedHand.Ankan)
	if err != nil {
		return nil, err
	}

	return &sortedHand, nil
}
