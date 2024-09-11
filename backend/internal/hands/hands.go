package hands

import (
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func GetHandScore(hand *types.PostHandScore) (*types.ReturnHandScore, error) {
	err := CheckValidData(hand)
	if err != nil {
		return nil, err
	}

	return nil, nil

}
