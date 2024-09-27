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
	var unsortedHand []string
	for _, block := range hand.Hand {
		for _, tile := range block {
			unsortedHand = append(unsortedHand, tile)
		}
	}
	return unsortedHand
}

func getHead(menzen [][]string) []string {
	for _, block := range menzen {
		if len(block) == 2 {
			return block
		}
	}
	return nil
}

func getMenzenKoutsuCount(menzen [][]string) int {
	var count int
	for _, block := range menzen {
		if checkValidKoutsu(block) {
			count++
		}
	}

	return count
}

func checkValidKoutsu(block []string) bool {
	if len(block) < 3 {
		return false
	}
	for i := 1; i < len(block); i++ {
		if block[i][:2] != block[i-1][:2] {
			return false
		}
	}

	return true
}
