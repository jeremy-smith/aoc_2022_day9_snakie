package main

import (
	"github.com/nsf/termbox-go"
	"math"
	"os"
	"sync"
	"time"
)

func handleInput(wg *sync.WaitGroup, dir chan<- string, exitChan chan<- int) {
	for {
		ev := termbox.PollEvent()
		if ev.Type != termbox.EventKey {
			continue
		}
		switch ev.Key {
		case termbox.KeyArrowUp:
			dir <- "U"
		case termbox.KeyArrowDown:
			dir <- "D"
		case termbox.KeyArrowLeft:
			dir <- "L"
		case termbox.KeyArrowRight:
			dir <- "R"
		case termbox.KeyEsc:
			wg.Done()
			exitChan <- 0
			return
		}

	}
}

func wormThings(wg *sync.WaitGroup, dirChan <-chan string, exitChan <-chan int) {
	ropeLen := 10
	knots := make([]pos, ropeLen)

	x, y := termbox.Size()

	// start positions
	for i := 0; i < ropeLen; i++ {
		knots[i] = pos{int64(math.Floor(float64(y) / 2)), int64(math.Floor(float64(x) / 2))}
	}

	var dir string

	for {
		select {
		case dir = <-dirChan:
		case <-exitChan:
			wg.Done()
			return
		default:
		}

		termbox.SetChar(int(knots[0].x), int(knots[0].y), ' ')
		for d := ropeLen - 2; d > 0; d-- {
			termbox.SetChar(int(knots[d].x), int(knots[d].y), ' ')
		}
		termbox.SetChar(int(knots[ropeLen-1].x), int(knots[ropeLen-1].y), ' ')

		knots[0] = move(knots[0], dir)

		for k := 1; k < ropeLen; k++ {
			if !areAdjacent(knots[k-1], knots[k]) {
				knots[k] = follow(knots[k-1], knots[k])
			}
		}

		termbox.SetChar(int(knots[0].x), int(knots[0].y), '@')
		for d := ropeLen - 2; d > 0; d-- {
			termbox.SetChar(int(knots[d].x), int(knots[d].y), '*')
		}
		termbox.SetChar(int(knots[ropeLen-1].x), int(knots[ropeLen-1].y), '¤')
		time.Sleep(100 * time.Millisecond)
		termbox.Flush()
	}
}

func main() {
	exitChan := make(chan int)

	termbox.Init()
	dir := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go handleInput(&wg, dir, exitChan)
	go wormThings(&wg, dir, exitChan)

	wg.Wait()
	termbox.Flush()
	termbox.Close()
	os.Exit(0)

}
