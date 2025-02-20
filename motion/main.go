package main

import "github.com/gonutz/prototype/draw"

func main(){
	title := "Text in motion"
	height, width := 350, 350

	i := 0

	draw.RunWindow(title, height, width, func(window draw.Window){
		i++
		window.FillEllipse(i%height, i%width, i%width, i%height, draw.Color{float32(i%7),float32(i%11), float32(i%17), 1})
	})
}