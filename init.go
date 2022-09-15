package main

import (
	"github.com/gdamore/tcell/v2"
)

func InitScreen() tcell.Screen {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := s.Init(); err != nil {
		panic(err)
	}

	s.EnableMouse()

	return s
}
