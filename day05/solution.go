package day05

import (
	"adventofcode2023/utils"
	"log"
	"math"
	"strconv"
	"strings"
)

type Almanac struct {
	Seeds                 []int
	SeedsToSoil           [][]int
	SoilToFertilizer      [][]int
	FertilizerToWater     [][]int
	WaterToLight          [][]int
	LightToTemperature    [][]int
	TemperatureToHumidity [][]int
	HumidityToLocation    [][]int
}

func (a *Almanac) getSeedLocation(seed int) int {

	iterable := [][][]int{
		a.SeedsToSoil,
		a.SoilToFertilizer,
		a.FertilizerToWater,
		a.WaterToLight,
		a.LightToTemperature,
		a.TemperatureToHumidity,
		a.HumidityToLocation,
	}

	value := seed
	for _, maps := range iterable {
		for i := 0; i < len(maps); i++ {
			if maps[i][1] <= value && value <= maps[i][1]+maps[i][2] {
				value = maps[i][0] + (value - maps[i][1])
				break
			}
		}
	}

	return value

}

type Range struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func parseInput(text string) Almanac {

	almanac := Almanac{
		Seeds:                 []int{},
		SeedsToSoil:           [][]int{},
		SoilToFertilizer:      [][]int{},
		FertilizerToWater:     [][]int{},
		WaterToLight:          [][]int{},
		LightToTemperature:    [][]int{},
		TemperatureToHumidity: [][]int{},
		HumidityToLocation:    [][]int{},
	}

	blocks := strings.Split(text, "\n\n")

	for _, block := range blocks {

		lines := strings.Split(block, "\n")
		label := strings.Fields(lines[0])[0]

		if label == "seeds:" {

			for _, s := range strings.Fields(strings.Split(lines[0], ":")[1]) {
				value, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				almanac.Seeds = append(almanac.Seeds, value)
			}

		} else {

			for _, line := range lines[1:] {
				switch label {
				case "seed-to-soil":
					almanac.SeedsToSoil = append(almanac.SeedsToSoil, parseMap(line))
					break
				case "soil-to-fertilizer":
					almanac.SoilToFertilizer = append(almanac.SoilToFertilizer, parseMap(line))
					break
				case "fertilizer-to-water":
					almanac.FertilizerToWater = append(almanac.FertilizerToWater, parseMap(line))
					break
				case "water-to-light":
					almanac.WaterToLight = append(almanac.WaterToLight, parseMap(line))
					break
				case "light-to-temperature":
					almanac.LightToTemperature = append(almanac.LightToTemperature, parseMap(line))
					break
				case "temperature-to-humidity":
					almanac.TemperatureToHumidity = append(almanac.TemperatureToHumidity, parseMap(line))
					break
				case "humidity-to-location":
					almanac.HumidityToLocation = append(almanac.HumidityToLocation, parseMap(line))
					break
				}
			}
		}

	}

	return almanac
}

func Solve() (string, string) {
	return PartOne(), PartTwo()
}

func PartOne() string {
	fr := utils.FileReader{Filepath: "./day05/input.txt"}
	almanac := parseInput(fr.ReadText())
	lowestLocation := math.MaxInt32
	for _, seed := range almanac.Seeds {
		location := almanac.getSeedLocation(seed)
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	return strconv.Itoa(lowestLocation)
}

func PartTwo() string {
	return ""
}
