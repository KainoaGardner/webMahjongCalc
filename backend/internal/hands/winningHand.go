package hands

import (
	"fmt"
	"sort"

	"github.com/KainoaGardner/webMahjongCalc/types"
)

func GetWinningHands(hand *types.HandParts) (*[]types.HandPartsBlocks, error) {
	// var winningHands []types.HandPartsBlocks

	//first menzen

	var possibleParts [][][]string
	tileCount := getTileMap(hand.Menzen)
	heads := getHeads(tileCount)
	mentsu := getMentsu(tileCount)

	return nil, nil
}

func getTileMap(handPart []string) map[string]int {
	tileCount := make(map[string]int)
	for _, tile := range handPart {
		tileCount[tile] += 1
	}

	return tileCount
}

func getHeads(tileCount map[string]int) [][]string {
	var heads [][]string
	for tile, count := range tileCount {
		if count >= 2 {
			heads = append(heads, []string{tile, tile})
		}
	}
	return heads

}

func getMentsu(tileCount map[string]int) [][]string {
	var mentsu [][]string
	for tile := range tileCount {
		shuntsu, ok := validShuntsu(tileCount, tile)
		if ok {
			mentsu = append(mentsu, shuntsu)
		}
		koutsu, ok := validKoutsu(tileCount, tile)
		if ok {
			mentsu = append(mentsu, koutsu)
		}
		kantsu, ok := validKantsu(tileCount, tile)
		if ok {
			mentsu = append(mentsu, kantsu)
		}

	}

	return mentsu
}

func validShuntsu(tileCount map[string]int, tile string) ([]string, bool) {
	suit := tile[0]
	if suit == 'H' {
		return nil, false
	}

	return nil, false
}

func validKoutsu(tileCount map[string]int, tile string) ([]string, bool) {
	if tileCount[tile] >= 3 {
		return []string{tile, tile, tile}, true
	} else if tileCount[tile] == 2 {
		if tile == "M5" && tileCount["M5A"] > 0 {
			return []string{tile, tile, "M5A"}, true
		} else if tile == "S5" && tileCount["S5A"] > 0 {
			return []string{tile, tile, "S5A"}, true
		} else if tile == "P5" && tileCount["P5A"] > 0 {
			return []string{tile, tile, "P5A"}, true
		}
	}

	return nil, false
}

func validKantsu(tileCount map[string]int, tile string) ([]string, bool) {
	if tileCount[tile] == 4 {
		return []string{tile, tile, tile, tile}, true
	} else if tileCount[tile] == 3 {
		if tile == "M5" && tileCount["M5A"] > 0 {
			return []string{tile, tile, tile, "M5A"}, true
		} else if tile == "S5" && tileCount["S5A"] > 0 {
			return []string{tile, tile, tile, "S5A"}, true
		} else if tile == "P5" && tileCount["P5A"] > 0 {
			return []string{tile, tile, tile, "P5A"}, true
		}
	}

	return nil, false
}
