package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

var (
	// Set default text style
	defStyle = tcell.StyleDefault.
				Background(tcell.ColorReset).
				Foreground(tcell.ColorReset)
	W int 
	H int
	Alive = make( map[cell]bool )
	Run = true 
)

func main() {
	fmt.Println("Conway's Game Of Life")
	
	s := InitScreen()

	W, H = s.Size()

	s.SetStyle(defStyle)
	go eventHandling(s)

	for {
		s.Clear()
		
		if Run {
			nGen := nextGen()

			// Erase Alive Content
			for c := range Alive { delete(Alive, c) }
		
			// Set new Alive 
			for c, v := range nGen {
				Alive[c] = v 
			}
		} else {
			for i, v := range "PAUSED" {
				s.SetContent(i, 0, v, nil, defStyle)
			}
		}

		// Update list of alive cells 
		// draw the cells on screen 
		for i := range Alive {
			s.SetContent(i.x, i.y, '+', nil, defStyle)
		}

		s.Show()
		time.Sleep(time.Millisecond * 50)
	}
}
