package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type Ball struct {
	x  int
	y  int
	dx int
	dy int
}

var (
	nballs    = 10
	contained = 0
	radius    = 100
	RAND_MAX  = 32767 // rand() in C provides a random number between 0 and something over 32767
)

func draw() {
	rand.Seed(time.Now().UnixNano())
	TermboxWidth, TermboxHeight := termbox.Size()
	maxX := TermboxWidth
	maxY := TermboxHeight * 2
	radius = (radius*radius + (maxX * maxY)) / 15000
	margin := 0
	if contained == 0 {
		margin = radius * 10
	}
	sumConst := 0.0225
	Balls := []Ball{} // array of balls
	for i := 0; i < nballs; i++ {
		x := rand.Intn(RAND_MAX)%(maxY-(2*margin)) + margin

		y := rand.Intn(RAND_MAX)%(maxY-2*margin) + margin
		dx := 0
		if rand.Intn(RAND_MAX)%2 == 0 {
			dx = -1
		} else {
			dx = 1
		}
		dy := 0
		if rand.Intn(RAND_MAX)%2 == 0 {
			dy = -1
		} else {
			dy = 1
		}
		Balls = append(Balls, Ball{x, y, dx, dy})
	}
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		for i := 0; i < nballs; i++ {
			if Balls[i].x+Balls[i].dx >= maxX-margin ||
				Balls[i].x+Balls[i].dx < margin {
				Balls[i].dx *= -1
			}
			if Balls[i].y+Balls[i].dy >= maxY-margin ||
				Balls[i].y+Balls[i].dy < margin {
				Balls[i].dy *= -1
			}
			Balls[i].x += Balls[i].dx
			Balls[i].y += Balls[i].dy

		}
		for i := 0; i < maxX; i++ {
			for j := 0; j < maxY; j++ {
				var sum [2]float64
				for j2 := 0; j2 < 2; j2++ { // using 2 as not custom
					for k := 0; k < nballs; k++ {
						y := j*2 + j2
						radiussquared := radius * radius
						sum[j2] += float64(radiussquared) / float64(((i-Balls[k].x)*(i-Balls[k].x) + (y-Balls[k].y)*(y-Balls[k].y)))
					}
				}
				if sum[0] > sumConst {
					if sum[1] > sumConst {
						// print i, j full block
						termbox.SetCell(i, j, '█', termbox.ColorCyan, termbox.ColorCyan)
					} else {
						termbox.SetCell(i, j, '▀', termbox.ColorCyan, termbox.ColorCyan)
					}
				} else if sum[1] > sumConst {
					termbox.SetCell(i, j, '▄', termbox.ColorCyan, termbox.ColorCyan)
				}
			}
		}
		time.Sleep(time.Duration(50000) * time.Microsecond)
		termbox.Flush()

	}
}

func main() {
	// Initialise termbox

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	event_queue := make(chan termbox.Event)

	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()
	go draw()
loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
		default:
			time.Sleep(5 * time.Second)
		}

	}
}
