package day04

import (
	"adventofcode2023/utils"
	"bufio"
	"log"
	"math"
	"strconv"
	"strings"
)

type ScratchCard struct {
	CardId         int64
	WinningNumbers []int64
	PlayedNumbers  []int64
}

func (sc *ScratchCard) GetValue() int64 {

	matchedNums := []int64{}
	for _, winningNum := range sc.WinningNumbers {
		for _, playedNum := range sc.PlayedNumbers {
			if winningNum == playedNum {
				matchedNums = append(matchedNums, winningNum)
				break
			}
		}
	}
	var value int64 = 0
	if len(matchedNums) > 0 {
		value = int64(math.Pow(float64(2), float64(len(matchedNums))-1))
	}
	return value
}

func (sc *ScratchCard) GetWinnings() []int64 {
	matchedNums := []int64{}
	for _, winningNum := range sc.WinningNumbers {
		for _, playedNum := range sc.PlayedNumbers {
			if winningNum == playedNum {
				matchedNums = append(matchedNums, winningNum)
				break
			}
		}
	}

	ids := []int64{}
	for i := sc.CardId + 1; i <= (sc.CardId + int64(len(matchedNums))); i++ {
		ids = append(ids, int64(i))
	}

	return ids
}

func parseInput(line string) ScratchCard {

	c := strings.Split(line, ": ")
	label := c[0]
	cardid, _ := strconv.ParseInt(strings.Fields(label)[1], 10, 64)
	nums := strings.Split(c[1], " | ")
	scanner := bufio.NewScanner(strings.NewReader(nums[0]))
	scanner.Split(bufio.ScanWords)

	winningNums := []int64{}
	for scanner.Scan() {
		winningNum, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		winningNums = append(winningNums, winningNum)
	}

	scanner = bufio.NewScanner(strings.NewReader(nums[1]))
	scanner.Split(bufio.ScanWords)

	playedNums := []int64{}
	for scanner.Scan() {
		playedNum, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		playedNums = append(playedNums, playedNum)
	}

	return ScratchCard{
		CardId:         cardid,
		WinningNumbers: winningNums,
		PlayedNumbers:  playedNums,
	}
}

func Solve() (string, string) {
	return PartOne(), PartTwo()
}

func PartOne() string {
	fr := utils.FileReader{Filepath: "./day04/input.txt"}
	lines := fr.ReadLines()

	var total int64 = 0
	for _, line := range lines {
		card := parseInput(line)
		total += card.GetValue()
	}

	return strconv.FormatInt(total, 10)
}

func PartTwo() string {
	fr := utils.FileReader{Filepath: "./day04/input.txt"}
	lines := fr.ReadLines()
	totals := map[int64]int64{}
	for i := 1; i <= len(lines); i++ {
		card := parseInput(lines[i-1])
		count := totals[card.CardId] + 1
		ids := card.GetWinnings()
		for _, id := range ids {
			totals[id] = totals[id] + count
		}
	}

	var total int64 = 0
	for _, value := range totals {
		total += value
	}

	return strconv.FormatInt(total, 10)
}
