package main

import (
	"github.com/gonutz/prototype/draw"
)

func main(){
	height := 350
	width := 550

	draw.RunWindow("Hello everyone", width, height, 
		func(window draw.Window){
			window.DrawText("Hello everyone", 10, 10, draw.Green)
		})
}