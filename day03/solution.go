package day03

import (
	"adventofcode2023/utils"
	"strconv"
	"strings"
)

type Schematic struct {
	Values [][]string
}

func (s *Schematic) FindPartNumbers() []PartNumber {

	numbers := []PartNumber{}

	for row := 0; row < len(s.Values); row++ {
		for col := 0; col < len(s.Values[row]); col++ {

			if strings.ContainsAny(s.Values[row][col], "0123456789") {

				partNumber := PartNumber{
					Points: []Point{{int64(row), int64(col)}},
					Value:  s.Values[row][col],
				}

				for i := col + 1; i < len(s.Values[row]); i, col = i+1, i {
					if strings.ContainsAny(s.Values[row][i], "0123456789") {
						partNumber.Points = append(partNumber.Points, Point{int64(row), int64(i)})
						partNumber.Value += s.Values[row][i]
					} else {
						break
					}
				}

				if s.CheckPartNumber(&partNumber) {
					numbers = append(numbers, partNumber)
				}

			}
		}
	}

	return numbers
}

func (s *Schematic) CheckPartNumber(pn *PartNumber) bool {

	var isValid bool = false

	for _, point := range pn.Points {

		adjacent := []Point{
			{point.y - 1, point.x - 1},
			{point.y - 1, point.x + 0},
			{point.y - 1, point.x + 1},
			{point.y + 0, point.x - 1},
			{point.y + 0, point.x + 1},
			{point.y + 1, point.x - 1},
			{point.y + 1, point.x + 0},
			{point.y + 1, point.x + 1},
		}

		for _, adj := range adjacent {
			value, err := s.GetPointValue(adj)
			if !err && !strings.ContainsAny(value, ".0123456789") {
				isValid = true
			}
		}
	}

	return isValid
}

func (s *Schematic) FindGears(partNumbers []PartNumber) []Gear {

	gears := []Gear{}

	for row := 0; row < len(s.Values); row++ {
		for col := 0; col < len(s.Values[row]); col++ {

			if s.Values[row][col] == "*" {

				matchingPartNumbers := []PartNumber{}
				point := Point{int64(row), int64(col)}

				adjacent := []Point{
					{point.y - 1, point.x - 1},
					{point.y - 1, point.x + 0},
					{point.y - 1, point.x + 1},
					{point.y + 0, point.x - 1},
					{point.y + 0, point.x + 1},
					{point.y + 1, point.x - 1},
					{point.y + 1, point.x + 0},
					{point.y + 1, point.x + 1},
				}

				for _, partNumber := range partNumbers {
					for _, pnp := range partNumber.Points {

						partNumberIsAdjacent := false
						for _, adj := range adjacent {
							if pnp == adj {
								partNumberIsAdjacent = true
								break
							}
						}

						if partNumberIsAdjacent {
							matchingPartNumbers = append(matchingPartNumbers, partNumber)
							break
						}

					}
				}

				if len(matchingPartNumbers) == 2 {
					gears = append(gears, Gear{
						Point:       point,
						PartNumbers: matchingPartNumbers,
					})
				}
			}

		}
	}

	return gears
}

func (s *Schematic) GetPointValue(p Point) (string, bool) {

	var value string = ""
	var err bool = false
	if p.y >= 0 && p.y < int64(len(s.Values)) && p.x >= 0 && p.x < int64(len(s.Values[0])) {
		value = s.Values[p.y][p.x]
	} else {
		err = true
	}
	return value, err
}

type PartNumber struct {
	Points []Point
	Value  string
}

type Gear struct {
	Point       Point
	PartNumbers []PartNumber
}

type Point struct {
	y int64
	x int64
}

func Solve() (string, string) {
	return PartOne(), PartTwo()
}

func PartOne() string {
	fr := utils.FileReader{Filepath: "./day03/input.txt"}
	lines := fr.ReadLines()

	data := [][]string{}
	for _, row := range lines {
		data = append(data, strings.Split(row, ""))
	}

	schematic := Schematic{data}
	partNumbers := schematic.FindPartNumbers()
	var total int64 = 0
	for _, pn := range partNumbers {
		value, _ := strconv.ParseInt(pn.Value, 10, 64)
		total += value
	}

	return strconv.FormatInt(total, 10)
}

func PartTwo() string {
	fr := utils.FileReader{Filepath: "./day03/input.txt"}
	lines := fr.ReadLines()

	data := [][]string{}
	for _, row := range lines {
		data = append(data, strings.Split(row, ""))
	}

	schematic := Schematic{data}
	partNumbers := schematic.FindPartNumbers()
	gears := schematic.FindGears(partNumbers)

	var total int64 = 0
	for _, gear := range gears {
		part1, _ := strconv.ParseInt(gear.PartNumbers[0].Value, 10, 64)
		part2, _ := strconv.ParseInt(gear.PartNumbers[1].Value, 10, 64)
		total += part1 * part2
	}

	return strconv.FormatInt(total, 10)
}
