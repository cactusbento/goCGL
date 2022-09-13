package main

import "github.com/gdamore/tcell/v2"

func eventHandling(s tcell.Screen) {
	for {
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				quit(s)
			}
		}
	}
}
