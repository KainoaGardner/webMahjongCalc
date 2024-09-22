package hands

import (
	"fmt"
	"internal/weak"
	"math"

	"github.com/KainoaGardner/webMahjongCalc/types"
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

		getHandScore(&currentHand)
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

func getHandScore(currentHand *types.WinningHand) {
	han, yakuList := GetYaku(currentHand)
	fu, fuList := GetFu(currentHand)
	score, scoreType, payments := getScore(han, fu, currentHand.ScoringParts)

	currentHand.HandScore.Han = han
	currentHand.HandScore.Fu = fu
	currentHand.HandScore.Score = score
	currentHand.HandScore.ScoreType = scoreType
	currentHand.HandScore.Payments = payments
	currentHand.HandScore.YakuList = yakuList
	currentHand.HandScore.FuList = fuList

}

func getScore(han int, fu int, scoringParts *types.HandScoringParts) (int, string, []int) {
	var handScore int

	basicPoints := fu * int(math.Pow(2, float64(2+han)))
	// honbaPoints += scoringParts.Honba * 300

	if scoringParts.Tsumo {

	}
	return 0, "", nil
}

// type Score struct {
// 	Han        int             `json:"han"`
// 	Fu         int             `json:"fu"`
// 	Score      int             `json:"score"`
// 	ScoreType  string          `json:"scoreType"`
// 	Yaku       []*YakuComponet `json:"yakuComponet"`
// 	FuComponet []*FuComponet   `json:"fuComponet"`
// }

// type WinningHand struct {
// 	Hand         [][]string        `json:"hand"`
// 	Open         bool              `json:"open"`
// 	ScoringParts *HandScoringParts `json:"scoringParts"`
// 	HandScore    *Score            `json:"handScore"`
// }
