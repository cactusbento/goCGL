package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	fmt.Println("Conway's Game Of Life")
	
	s := InitScreen()

	// Set default text style
	DefStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(DefStyle)

	x := 0

	
	go eventHandling(s)

	for {
		s.Clear()


		s.SetContent(x, 0, '=', nil, DefStyle)
		x++

		if w, _ := s.Size(); x >= w - 1 {
			quit(s)
		}

		s.Show()
		time.Sleep(time.Millisecond * 100)
	}
}
