package main

import (
	"log"

	"github.com/gen2brain/raylib-go/raylib"
)

var (
	Mkill = false 
	W int 
	H int 
	Size = 10
	Alive = make( map[cell]bool )
	Run = true 
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 600, "Conway's Game of Life ")
	rl.SetTargetFPS(30)


	log.Println("Game Started.")

	for !rl.WindowShouldClose() {
		W, H = rl.GetScreenWidth(), rl.GetScreenHeight()

		W /= Size
		H /= Size

		rl.BeginDrawing()

		eventHandler()

		nGen := nextGen()
		
		if Run {
			// Erase Alive Content
			for c := range Alive { delete(Alive, c) }
		
			// Set new Alive 
			for c, v := range nGen {
				Alive[c] = v 
			}
		} else {
			rl.DrawText("PAUSED", 0, 0, 20, rl.LightGray)
		}

		rl.ClearBackground(rl.DarkGray)

		// draw the cells on screen 
		for i := range Alive {
			rec := rl.Rectangle{
				X: float32(i.x * Size),
				Y: float32(i.y * Size), 
				Width: float32(Size), 
				Height: float32(Size), 
			}

			rl.DrawRectangleLinesEx(rec, 2.0, rl.White)
		}


		rl.EndDrawing()
	}
	rl.CloseWindow()
}

