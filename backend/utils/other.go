package utils

import (
	"github.com/KainoaGardner/webMahjongCalc/types"
)

func GetTileMap(handPart []string) map[string]int {
	tileCount := make(map[string]int)
	AddToTileMap(tileCount, handPart)
	return tileCount
}

// add to tileMap but only first two letters no aka
func AddToTileMap(tileMap map[string]int, handPart []string) {
	for _, tile := range handPart {
		tileMap[tile[:2]] += 1
	}
}

func CheckOpenHand(hand *types.WinningHand) bool {
	if len(hand.HandParts.Chi) > 0 || len(hand.HandParts.Pon) > 0 || len(hand.HandParts.Kan) > 0 {
		return true
	}
	return false
}

func GetUnsortedHand(hand *types.WinningHand) []string {
	var unsortedHand []string
	for _, block := range hand.Hand {
		for _, tile := range block {
			unsortedHand = append(unsortedHand, tile)
		}
	}
	return unsortedHand
}

func GetHead(menzen [][]string) ([]string, bool) {
	for _, block := range menzen {
		if len(block) == 2 {
			return block, true
		}
	}
	return nil, false
}

func GetMenzenKoutsuCount(menzen [][]string) int {
	var count int
	for _, block := range menzen {
		if CheckKoutsuBlock(block) {
			count++
		}
	}

	return count
}

func CheckKoutsuBlock(block []string) bool {
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

func CheckShuntsuBlock(block []string) bool {
	if len(block) != 3 {
		return false
	}
	startNumber := int(block[0][1])

	for i := 1; i < len(block); i++ {
		if int(block[i][1]) != startNumber+i {
			return false
		}
	}

	return true
}

func GetBlockString(block []string) string {
	var result string
	suit := block[0][0]
	result += string(suit)

	for _, tile := range block {
		result += string(tile[1])
	}

	return result
}

func RemoveYaku(hand []*types.YakuComponet, yaku string) []*types.YakuComponet {
	for i := len(hand) - 1; i >= 0; i-- {
		if hand[i].Title == yaku {
			return append(hand[:i], hand[i+1:]...)

		}

	}
	return hand

}

func CheckYaku(hand []*types.YakuComponet, yakuTitle string) bool {
	for _, yaku := range hand {
		if yaku.Title == yakuTitle {
			return true
		}

	}
	return false

}
