package hands

import (
	"fmt"
	"github.com/KainoaGardner/webMahjongCalc/types"
	"maps"
	"strings"
)

func getValidHands(hand *types.HandParts) ([]types.WinningHand, error) {
	//first menzen
	totalTileCount := make(map[string]int)
	var potentialHands [][][]string

	//mentsu
	tileCount := getTileMap(hand.Menzen)
	addToTileMap(totalTileCount, hand.Menzen)

	//check kokushi
	kokushiHand, ok := checkKokushimusou(tileCount)
	if ok {
		potentialHands = append(potentialHands, kokushiHand)
	}

	//check chiitoitsu
	chiitoitsuHand, ok := checkChiitoitsu(tileCount)
	if ok {
		potentialHands = append(potentialHands, chiitoitsuHand)
	}

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
	potentialHands, err := getPotentialHands(potentialHands, heads, mentsu)
	if err != nil {
		return nil, err
	}

	validHands := filterPotentialHands(potentialHands, totalTileCount)

	//add akadora back

	formattedHands, err := formatHands(validHands, hand)
	if err != nil {
		return nil, err
	}

	akaDora := getAllAkaDora(*hand)
	handParts := [][]string{hand.Menzen, hand.Chi, hand.Pon, hand.Kan, hand.Ankan}
	err = replaceAllAkadora(akaDora, formattedHands, handParts)
	if err != nil {
		return nil, err
	}

	return formattedHands, nil
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

func getPotentialHands(potentialHands [][][]string, heads [][]string, mentsu [][]string) ([][][]string, error) {
	sortMentsu(mentsu)

	for i := 0; i < len(heads); i++ {
		_, potentialHands = btPotentialHands(0, mentsu, [][]string{heads[i]}, potentialHands)
	}
	if len(potentialHands) == 0 {
		return nil, fmt.Errorf("No vaild potentialhands")
	}
	return potentialHands, nil
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

func checkKokushimusou(tileCount map[string]int) ([][]string, bool) {
	kokushi := []string{"H1", "H2", "H3", "H4", "H5", "H6", "H7", "M1", "M9", "S1", "S9", "P1", "P9"}
	var kokushiHand [][]string

	//check at least one of each honor terminal
	for _, tile := range kokushi {
		if tileCount[tile] < 1 {
			return nil, false
		}

		kokushiHand = append(kokushiHand, []string{tile})
	}

	//check at least 1 has more than 1
	for i, tile := range kokushi {
		if tileCount[tile] > 1 {
			kokushiHand[i] = append(kokushiHand[i], tile)
			return kokushiHand, true
		}
	}

	return nil, false
}

func checkChiitoitsu(tileCount map[string]int) ([][]string, bool) {
	var chiitoitsuHand [][]string
	for tile, count := range tileCount {
		if count != 2 {
			return nil, false
		}
		chiitoitsuHand = append(chiitoitsuHand, []string{tile, tile})
	}
	return chiitoitsuHand, true
}

func formatHands(validHands [][][]string, hand *types.HandParts) ([]types.WinningHand, error) {
	formattedHands := []types.WinningHand{}

	for _, validHand := range validHands {
		formattedHand := types.WinningHand{
			Hand:      validHand,
			HandParts: &types.HandPartBlocks{},
			HandScore: &types.Score{},
		}
		formattedHand.HandParts.Agari = hand.Menzen[len(hand.Menzen)-1]

		menTileCount := getTileMap(hand.Menzen)
		var menzen [][]string
		chiTileCount := getTileMap(hand.Chi)
		var chi [][]string
		ponTileCount := getTileMap(hand.Pon)
		var pon [][]string
		kanTileCount := getTileMap(hand.Kan)
		var kan [][]string
		ankanTileCount := getTileMap(hand.Ankan)
		var ankan [][]string

		tileCounts := []map[string]int{menTileCount, chiTileCount, ponTileCount, kanTileCount, ankanTileCount}

		handParts := [][][]string{menzen, chi, pon, kan, ankan}

		for i := len(validHand) - 1; i >= 0; i-- {
			handParts = addBlockToParts(validHand[i], tileCounts, handParts)

		}

		formattedHand.HandParts.Menzen = handParts[0] //menzen
		formattedHand.HandParts.Chi = handParts[1]    //chi
		formattedHand.HandParts.Pon = handParts[2]    //pon
		formattedHand.HandParts.Kan = handParts[3]    //kan
		formattedHand.HandParts.Ankan = handParts[4]  //ankan

		formattedHands = append(formattedHands, formattedHand)

	}

	return formattedHands, nil
}

func addBlockToParts(block []string, tileCounts []map[string]int, handParts [][][]string) [][][]string {
	blockTileCount := getTileMap(block)

	for i := 0; i < len(handParts); i++ {
		if checkAddBlockToPart(blockTileCount, tileCounts[i]) {
			handParts[i] = append(handParts[i], block)
			for tile, amount := range blockTileCount {
				tileCounts[i][tile] -= amount
			}
			break
		}
	}

	return handParts
}

func checkAddBlockToPart(blockTileCount map[string]int, tileCount map[string]int) bool {
	for tile, amount := range blockTileCount {
		if tileCount[tile] < amount {
			return false
		}
	}

	return true
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

func replaceAllAkadora(akaDoraList []string, winningHands []types.WinningHand, handParts [][]string) error {
	for _, akaDora := range akaDoraList {
		akaDoraPart, err := getAkaDoraPart(akaDora, handParts)
		if err != nil {
			return err
		}
		for _, winningHand := range winningHands {
			switch akaDoraPart {
			case 0:
				replaceAkaDora(akaDora, winningHand.HandParts.Menzen)
			case 1:
				replaceAkaDora(akaDora, winningHand.HandParts.Chi)
			case 2:
				replaceAkaDora(akaDora, winningHand.HandParts.Pon)
			case 3:
				replaceAkaDora(akaDora, winningHand.HandParts.Kan)
			case 4:
				replaceAkaDora(akaDora, winningHand.HandParts.Ankan)
			default:
				return fmt.Errorf("Cant find akadora part for %s", akaDora)

			}

		}

	}
	return nil
}

func getAkaDoraPart(akaDora string, handParts [][]string) (int, error) {
	for i, part := range handParts {
		for _, tile := range part {
			if tile == akaDora {
				return i, nil
			}
		}

	}

	return 0, fmt.Errorf("Cant find akadora part for %s", akaDora)

}

func replaceAkaDora(dora string, handPart [][]string) {
	for i, block := range handPart {
		for j, tile := range block {
			if tile == dora[:2] {
				handPart[i][j] = dora

				return
			}
		}
	}

}
