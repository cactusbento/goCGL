package main

import "github.com/gdamore/tcell/v2"

func eventHandling(s tcell.Screen) {
	for {
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			W, H = ev.Size()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC || ev.Key() == tcell.KeyEscape {
				quit(s)
			}
		case *tcell.EventMouse:
			if ev.Buttons() == tcell.ButtonPrimary {
				posX, posY := ev.Position()
				posCell := cell{x: posX, y: posY}

				if isAlive(posCell) {
					delete(Alive, posCell)
				} else {
					Alive[ cell{x: posX, y: posY} ] = true
				}
			} 

			if ev.Buttons() == tcell.ButtonSecondary {
				Run = !Run 
			}
		}
	}
}
