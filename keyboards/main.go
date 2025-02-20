package main

import "github.com/gonutz/prototype/draw"

func main(){
	title := "Text in motion"
	height, width := 300, 300

	CurrentX := width / 2
	CurrentY := height / 2

	draw.RunWindow(title, height, width, func(window draw.Window){
		if window.IsKeyDown(draw.KeyDown){
			CurrentY += 1
		} else if window.IsKeyDown(draw.KeyUp){
			CurrentY -= 1
		} else if window.IsKeyDown(draw.KeyLeft){
			CurrentX -= 1
		} else if window.IsKeyDown(draw.KeyRight){
			CurrentX += 1
		}

		window.FillRect(CurrentX, CurrentY, 50, 60, draw.Blue)
	})
}