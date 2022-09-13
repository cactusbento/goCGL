package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	fmt.Println("Conway's Game Of Life")
	
	s := InitScreen()

	// Set default text style
	DefStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(DefStyle)

	quit := func() {
		s.Fini()
		os.Exit(0)
	}

	x := 0

	
	go func() {
		for {
			ev := s.PollEvent()

			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyCtrlC {
					quit()
				}
			}
		}
	}()

	for {
		s.Clear()


		s.SetContent(x, 0, '=', nil, DefStyle)
		x++

		if w, _ := s.Size(); x >= w - 1 {
			quit()
		}

		s.Show()
		time.Sleep(time.Millisecond * 100)
	}
}
