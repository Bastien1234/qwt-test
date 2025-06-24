package cmd

import (
	"fmt"
	gameparser "qwant/internal/parser"
	"sync"
)

const (
	file = "inputs/input.txt"
)

func Main() int {
	game, err := gameparser.InitGameFromInput(file)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	var wg sync.WaitGroup
	for _, mower := range game.Mowers {
		wg.Add(1)
		func() {
			defer wg.Done()

			mower.ExecuteMoving(game.Lawn)
		}()
	}
	wg.Wait()

	for _, mower := range game.Mowers {
		mower.PrintFinalPosition()
	}
	return 0
}
