package gameparser

import (
	"bufio"
	"fmt"
	"os"
	"qwant/internal/core"
	"strconv"
	"strings"
)

type Game struct {
	Lawn   *core.Lawn
	Mowers []*core.Mower
}

func InitGameFromInput(filePath string) (*Game, error) {
	game := &Game{}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return nil, fmt.Errorf("missing lawn dimensions")
	}
	dimensionLine := strings.TrimSpace(scanner.Text())
	dimensionLine = removeBOM(dimensionLine)
	dimensions := strings.Fields(dimensionLine)
	if len(dimensions) != 2 {
		return nil, fmt.Errorf("invalid lawn dimensions format")
	}

	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return nil, fmt.Errorf("invalid lawn width: %v", err)
	}

	height, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return nil, fmt.Errorf("invalid lawn height: %v", err)
	}

	// Init lawn
	lawn := core.InitLawn(width, height)
	game.Lawn = lawn

	newMowerId := int8(1)
	for scanner.Scan() {
		positionLine := strings.TrimSpace(scanner.Text())
		if positionLine == "" {
			continue
		}

		positionFields := strings.Fields(positionLine)
		if len(positionFields) != 3 {
			return nil, fmt.Errorf("invalid mower position format: %s", positionLine)
		}

		x, err := strconv.Atoi(positionFields[0])
		if err != nil {
			return nil, fmt.Errorf("invalid x coordinate: %v", err)
		}

		y, err := strconv.Atoi(positionFields[1])
		if err != nil {
			return nil, fmt.Errorf("invalid y coordinate: %v", err)
		}

		direction, err := core.ParseDirection(positionFields[2])
		if err != nil {
			return nil, fmt.Errorf("invalid direction: %v", err)
		}

		// Read mower instructions
		if !scanner.Scan() {
			return nil, fmt.Errorf("missing instructions for mower at (%d,%d)", x, y)
		}

		actions := strings.TrimSpace(scanner.Text())

		// Init new mower
		position := core.Position{X: x, Y: y}
		mower := core.InitMower(newMowerId, position, direction, actions)
		game.Mowers = append(game.Mowers, mower)
		newMowerId++
	}

	return game, nil
}

func removeBOM(s string) string {
	if strings.HasPrefix(s, "\uFEFF") {
		return strings.TrimPrefix(s, "\uFEFF")
	}
	return s
}
