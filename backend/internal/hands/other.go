package hands

import (
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func getTileMap(handPart []string) map[string]int {
	tileCount := make(map[string]int)
	addToTileMap(tileCount, handPart)
	return tileCount
}

// add to tileMap but only first two letters no aka
func addToTileMap(tileMap map[string]int, handPart []string) {
	for _, tile := range handPart {
		tileMap[tile[:2]] += 1
	}
}

func checkOpenHand(hand *types.WinningHand) bool {
	if len(hand.HandParts.Chi) > 0 || len(hand.HandParts.Pon) > 0 || len(hand.HandParts.Kan) > 0 {
		return true
	}
	return false
}

func getUnsortedHand(hand *types.WinningHand) []string {
	return nil
}
