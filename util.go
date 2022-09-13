package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

func quit(s tcell.Screen) {
	s.Fini()
	os.Exit(0)
}
