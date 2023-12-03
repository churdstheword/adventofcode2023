package day02

import (
	"adventofcode2023/utils"
	"strconv"
	"strings"
)

type Game struct {
	GameId int64
	Rules  GameRules
	Rounds []GameRound
}

func (game *Game) validate() bool {
	isValid := true
	for _, round := range game.Rounds {
		if !round.validate(&game.Rules) {
			isValid = false
		}
	}
	return isValid
}

func (game *Game) getPower() int64 {
	var rMax, gMax, bMax int64 = 0, 0, 0
	for _, round := range game.Rounds {
		if round.r > rMax {
			rMax = round.r
		}
		if round.g > gMax {
			gMax = round.g
		}
		if round.b > bMax {
			bMax = round.b
		}
	}

	return rMax * gMax * bMax
}

type GameRound struct {
	r int64
	g int64
	b int64
}

func (round *GameRound) validate(rules *GameRules) bool {
	return (round.r <= rules.TotalRedCubes && round.g <= rules.TotalGreenCubes && round.b <= rules.TotalBlueCubes)
}

type GameRules struct {
	TotalRedCubes   int64
	TotalGreenCubes int64
	TotalBlueCubes  int64
}

func Solve() (string, string) {
	return PartOne(), PartTwo()
}

func PartOne() string {
	fr := utils.FileReader{Filepath: "./day02/input.txt"}
	lines := fr.ReadLines()
	var total int64 = 0
	for _, line := range lines {
		game := parseGame(line)
		if game.validate() {
			total += game.GameId
		}
	}
	return strconv.FormatInt(total, 10)
}

func PartTwo() string {
	fr := utils.FileReader{Filepath: "./day02/input.txt"}
	lines := fr.ReadLines()
	var total int64 = 0
	for _, line := range lines {
		game := parseGame(line)
		total += game.getPower()
	}

	return strconv.FormatInt(total, 10)
}

func parseGame(line string) Game {
	gameData := strings.Split(line, ": ")
	gameid, _ := strconv.ParseInt(gameData[0][5:], 10, 64)
	gameRounds := []GameRound{}
	rounds := strings.Split(gameData[1], "; ")
	for _, round := range rounds {
		values := strings.Split(round, ", ")
		var gameRound = GameRound{0, 0, 0}
		for _, value := range values {
			cube := strings.Split(value, " ")
			switch label := cube[1]; label {
			case "red":
				gameRound.r, _ = strconv.ParseInt(cube[0], 10, 64)
			case "green":
				gameRound.g, _ = strconv.ParseInt(cube[0], 10, 64)
			case "blue":
				gameRound.b, _ = strconv.ParseInt(cube[0], 10, 64)
			}
		}
		gameRounds = append(gameRounds, gameRound)
	}

	return Game{
		GameId: gameid,
		Rules:  GameRules{12, 13, 14},
		Rounds: gameRounds,
	}
}
