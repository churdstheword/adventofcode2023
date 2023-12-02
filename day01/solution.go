package day01

import (
	"adventofcode2023/utils"
	"strconv"
	"strings"
)

var numbersMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Solve() (string, string) {
	return PartOne(), PartTwo()
}

func PartOne() string {

	fr := utils.FileReader{Filepath: "./day01/input.txt"}
	lines := fr.ReadLines()

	var total int = 0
	for _, line := range lines {
		total = total + calibrate(line)
	}

	return strconv.Itoa(total)

}

func PartTwo() string {
	fr := utils.FileReader{Filepath: "./day01/input.txt"}
	lines := fr.ReadLines()

	var total int = 0
	for _, line := range lines {
		total = total + calibrateV2(line)
	}

	return strconv.Itoa(total)
}

func calibrate(line string) int {
	s := strings.Split(line, "")
	first := s[strings.IndexAny(line, "0123456789")]
	last := s[strings.LastIndexAny(line, "0123456789")]
	value, _ := strconv.Atoi(first + last)
	return value
}

func calibrateV2(line string) int {

	var first string = "0"
	var last string = "0"

	runes := strings.Split(line, "")

left:
	for left := 0; left < len(runes); left++ {
		substr := strings.Join(runes[left:], "")
		for k, v := range numbersMap {
			if strings.HasPrefix(substr, k) || strings.HasPrefix(substr, v) {
				first = v
				break left
			}
		}
	}

right:
	for right := len(runes); right >= 0; right-- {
		substr := strings.Join(runes[:right], "")
		for k, v := range numbersMap {
			if strings.HasSuffix(substr, k) || strings.HasSuffix(substr, v) {
				last = v
				break right
			}
		}
	}

	value, _ := strconv.Atoi(first + last)
	return value

}
