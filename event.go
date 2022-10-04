package main

import (
	"log"

	"github.com/gen2brain/raylib-go/raylib"
)

func eventHandler() {
	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		Run = !Run
		if !Run { 
			log.Println("PAUSED") 
		} else { log.Println("RESUMED")}
	}

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		mPos := rl.GetMousePosition()
		posCell := cell{ x: int(mPos.X/float32(Size)),
						 y: int(mPos.Y/float32(Size))}
		
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if isAlive(posCell) {
				log.Println("Manual Kill:", posCell)
				Mkill = true 
				delete(Alive, posCell) 
			} else {
				log.Println("Manual Live:", posCell)
				Mkill = false 
				Alive[ posCell ] = true 
			}
		} else { 
			if !Run {
				posCell := cell{ x: int(mPos.X/float32(Size)),
								 y: int(mPos.Y/float32(Size))}


				if isAlive(posCell) && Mkill {
					log.Println("Manual Kill:", posCell)
					delete(Alive, posCell)
				}
				if !isAlive(posCell) && !Mkill {
					log.Println("Manual Live:", posCell)
					Alive[ posCell ] = true 
				}
			}
		}
	}
}

