package main

import (
	"log"
	"os"
	"time"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

type void struct{}

// x, y = center
func drawText(s tcell.Screen, x, y int, text []string) {
	x0 := x - utf8.RuneCountInString(text[0])/2
	y0 := y - len(text)/2
	for i := range text {
		row := []rune(text[i])
		for j := range row {
			s.SetContent(x0+j, y0+i, row[j], nil, tcell.StyleDefault)
		}
	}
}

func drawChopsticks(s tcell.Screen, chopsticks []chopstick) {
	w, h := s.Size()
	if chopsticks[nw].owner == nil {
		s.SetContent(7, 4, '╲', nil, tcell.StyleDefault)
		s.SetContent(8, 5, '╲', nil, tcell.StyleDefault)
	}
	if chopsticks[ne].owner == nil {
		s.SetContent(w-8, 4, '╱', nil, tcell.StyleDefault)
		s.SetContent(w-9, 5, '╱', nil, tcell.StyleDefault)
	}
	if chopsticks[se].owner == nil {
		s.SetContent(w-8, h-4, '╲', nil, tcell.StyleDefault)
		s.SetContent(w-9, h-5, '╲', nil, tcell.StyleDefault)
	}
	if chopsticks[sw].owner == nil {
		s.SetContent(7, h-4, '╱', nil, tcell.StyleDefault)
		s.SetContent(8, h-5, '╱', nil, tcell.StyleDefault)
	}
}

func drawPhilosophers(s tcell.Screen, philosophers []philosopher) {
	w, h := s.Size()
	drawText(s, w/2, 4, philosophers[north].ascii())
	drawText(s, w-5, (h+2-1)/2, philosophers[east].ascii())
	drawText(s, w/2, h-5, philosophers[south].ascii())
	drawText(s, 5, (h+2-1)/2, philosophers[west].ascii())
}

func draw(s tcell.Screen, philosophers []philosopher, chopsticks []chopstick) {
	s.Clear()
	drawPhilosophers(s, philosophers)
	drawChopsticks(s, chopsticks)
	s.Sync()
}

func philosopherThread(start chan void, done chan bool, p *philosopher) {
	for {
		<-start
		if p.eating {
			p.eating = false
			p.hasEaten = true
			done <- true
		} else if p.hasEaten {
			p.hasEaten = false
			p.left.owner = nil
			p.right.owner = nil
			done <- true
		} else if p.left.owner == p && p.right.owner == p {
			p.eating = true
			done <- true
		} else if p.firstStick().owner == p {
			done <- p.pickUp(p.secondStick())
		} else {
			done <- p.pickUp(p.firstStick())
		}
	}
}

func main() {
	encoding.Register()

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Init(); err != nil {
		log.Fatal(err)
	}

	chopsticks := []chopstick{
		chopstick{number: 0},
		chopstick{number: 1},
		chopstick{number: 2},
		chopstick{number: 3},
	}

	philosophers := []philosopher{
		philosopher{north, &chopsticks[ne], &chopsticks[nw], false, false},
		philosopher{east, &chopsticks[se], &chopsticks[ne], false, false},
		philosopher{south, &chopsticks[sw], &chopsticks[se], false, false},
		philosopher{west, &chopsticks[nw], &chopsticks[sw], false, false},
	}

	draw(s, philosophers, chopsticks)

	start := make(chan void)
	done := make(chan bool)

	for i := range philosophers {
		go philosopherThread(start, done, &philosophers[i])
	}

	go func() {
		for range time.Tick(700 * time.Millisecond) {
			for {
				start <- void{}
				if <-done {
					break
				}
			}
			draw(s, philosophers, chopsticks)
		}
	}()

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			draw(s, philosophers, chopsticks)
		case *tcell.EventKey:
			if ev.Rune() == 'q' {
				s.Fini()
				os.Exit(0)
			}
		}
	}
}
