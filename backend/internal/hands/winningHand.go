package hands

import (
	"fmt"
	"maps"
	"strings"

	"github.com/KainoaGardner/webMahjongCalc/types"
)

func GetValidHands(hand *types.HandParts) ([][][]string, error) {

	//check kokushi

	//check chiitoitsu

	//first menzen
	totalTileCount := make(map[string]int)

	akaDora := getAllAkaDora(*hand)

	//mentsu
	tileCount := getTileMap(hand.Menzen)
	addToTileMap(totalTileCount, hand.Menzen)

	heads := getHeads(tileCount)
	mentsu := getMentsu(tileCount)

	//calls
	tileCount = getTileMap(hand.Chi)
	addToTileMap(totalTileCount, hand.Chi)
	mentsu = appendChiMentsu(tileCount, mentsu)

	tileCount = getTileMap(hand.Pon)
	addToTileMap(totalTileCount, hand.Pon)
	mentsu = appendPonMentsu(tileCount, mentsu)

	tileCount = getTileMap(hand.Kan)
	addToTileMap(totalTileCount, hand.Kan)
	mentsu = appendKanMentsu(tileCount, mentsu)

	tileCount = getTileMap(hand.Ankan)
	addToTileMap(totalTileCount, hand.Ankan)
	mentsu = appendKanMentsu(tileCount, mentsu)

	//check valid hands
	potentialHands, err := getPotentialHands(heads, mentsu)
	if err != nil {
		return nil, err
	}

	validHands := filterPotentialHands(potentialHands, totalTileCount)

	//add akadora back
	replaceAllAkadora(akaDora, validHands)

	return validHands, nil
}

func addToTileMap(tileMap map[string]int, handPart []string) {
	for _, tile := range handPart {
		tileMap[tile[:2]] += 1
	}
}

func getTileMap(handPart []string) map[string]int {
	tileCount := make(map[string]int)
	addToTileMap(tileCount, handPart)
	return tileCount
}

func getAllAkaDora(hand types.HandParts) []string {
	var akaDora []string
	akaDora = getAkaDora(akaDora, hand.Menzen)
	akaDora = getAkaDora(akaDora, hand.Chi)
	akaDora = getAkaDora(akaDora, hand.Pon)
	akaDora = getAkaDora(akaDora, hand.Kan)
	akaDora = getAkaDora(akaDora, hand.Ankan)

	return akaDora
}

func getAkaDora(akaDora []string, part []string) []string {
	for _, tile := range part {
		if tile == "M5A" || tile == "S5A" || tile == "P5A" {
			akaDora = append(akaDora, tile)
		}
	}
	return akaDora
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
			for _, set := range shuntsu {
				mentsu = append(mentsu, set)
			}
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

func validShuntsu(tileCount map[string]int, tile string) ([][]string, bool) {
	suit := tile[0]
	if suit == 'H' {
		return nil, false
	}

	var result [][]string
	var shuntsu []string

	startNumber := int(tile[1])
	minNumberCount := tileCount[tile]
	for i := 0; i < 3; i++ {
		nextNumber := fmt.Sprintf("%c%c", suit, startNumber+i)

		if tileCount[nextNumber] > 0 {
			shuntsu = append(shuntsu, nextNumber)
			if tileCount[nextNumber] < minNumberCount {
				minNumberCount = tileCount[nextNumber]
			}
		} else {
			return nil, false
		}
	}
	for i := 0; i < minNumberCount; i++ {
		result = append(result, shuntsu)
	}

	return result, true
}

func validKoutsu(tileCount map[string]int, tile string) ([]string, bool) {
	if tileCount[tile] >= 3 {
		return []string{tile, tile, tile}, true
	}

	return nil, false
}

func validKantsu(tileCount map[string]int, tile string) ([]string, bool) {
	if tileCount[tile] == 4 {
		return []string{tile, tile, tile, tile}, true
	}

	return nil, false
}

func appendChiMentsu(tileCount map[string]int, mentsu [][]string) [][]string {
	for tile := range tileCount {
		shuntsu, ok := validShuntsu(tileCount, tile)
		if ok {
			for _, set := range shuntsu {
				mentsu = append(mentsu, set)
			}
		}
	}
	return mentsu
}

func appendPonMentsu(tileCount map[string]int, mentsu [][]string) [][]string {
	for tile := range tileCount {
		koutsu, ok := validKoutsu(tileCount, tile)
		if ok {
			mentsu = append(mentsu, koutsu)
		}
	}
	return mentsu
}

func appendKanMentsu(tileCount map[string]int, mentsu [][]string) [][]string {
	for tile := range tileCount {
		kantsu, ok := validKantsu(tileCount, tile)
		if ok {
			mentsu = append(mentsu, kantsu)
		}
	}
	return mentsu
}

func getPotentialHands(heads [][]string, mentsu [][]string) ([][][]string, error) {
	var result [][][]string
	sortMentsu(mentsu)

	for i := 0; i < len(heads); i++ {
		_, result = btPotentialHands(0, mentsu, [][]string{heads[i]}, result)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("No vaild potentialhands")
	}
	return result, nil
}

func btPotentialHands(index int, mentsu [][]string, hand [][]string, result [][][]string) ([][]string, [][][]string) {
	if len(hand) == 5 {
		temp := make([][]string, len(hand))
		copy(temp, hand)
		result = append(result, temp)
		return hand, result
	}

	var prev []string
	for i := index; i < len(mentsu); i++ {
		if compareMentsu(mentsu[i], prev) == 0 {
			continue
		}

		temp := make([]string, len(mentsu[i]))
		copy(temp, mentsu[i])
		hand = append(hand, temp)
		hand, result = btPotentialHands(i+1, mentsu, hand, result)
		hand = hand[:len(hand)-1]
		prev = mentsu[i]

	}

	return hand, result
}

// insertion sort
func sortMentsu(mentsu [][]string) {
	for i := 1; i < len(mentsu); i++ {
		key := mentsu[i]
		j := i - 1

		for j >= 0 && compareMentsu(mentsu[j], key) == 1 {
			mentsu[j+1] = mentsu[j]
			j--
		}
		mentsu[j+1] = key
	}
}

// a == b = 0, a > b = 1, a < b = -1
func compareMentsu(a []string, b []string) int {
	var aTotal, bTotal string
	for _, tile := range a {
		aTotal += tile
	}
	for _, tile := range b {
		bTotal += tile
	}

	return strings.Compare(aTotal, bTotal)
}

func checkPotentialHandTileAmount(potentialHand [][]string, totalTileCount map[string]int) bool {
	for _, block := range potentialHand {
		for _, tile := range block {
			totalTileCount[tile] -= 1
		}
	}

	for _, count := range totalTileCount {
		if count != 0 {
			return false
		}
	}

	return true
}

func filterPotentialHands(potentialHands [][][]string, totalTileCount map[string]int) [][][]string {

	for i := len(potentialHands) - 1; i >= 0; i-- {

		tileCountCopy := maps.Clone(totalTileCount)
		if !checkPotentialHandTileAmount(potentialHands[i], tileCountCopy) {
			potentialHands = removeHandByIndex(potentialHands, i)
		}
	}
	return potentialHands

}

func removeHandByIndex(potentialHands [][][]string, index int) [][][]string {
	potentialHands[index] = potentialHands[len(potentialHands)-1]
	return potentialHands[:len(potentialHands)-1]
}

func replaceAllAkadora(akaDora []string, validHands [][][]string) {
	for _, dora := range akaDora {
		for _, validHand := range validHands {
			replaceAkaDora(dora, validHand)
		}
	}
}

func replaceAkaDora(dora string, validHand [][]string) {
	for i, block := range validHand {
		for j, tile := range block {
			if tile == dora[:2] {
				validHand[i][j] = dora

				return
			}
		}
	}

}
