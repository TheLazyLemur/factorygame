package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func snapToGrid(pos rl.Vector2) rl.Vector2 {
	x := float32(math.Floor(float64(pos.X/50)) * 50)
	y := float32(math.Floor(float64(pos.Y/50)) * 50)

	return rl.NewVector2(x, y)
}

func getWorldMousePos(camera rl.Camera2D) rl.Vector2 {
	mousePos := rl.GetMousePosition()
	return rl.GetScreenToWorld2D(mousePos, camera)
}
