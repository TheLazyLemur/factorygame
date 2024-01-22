package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Building struct {
	Pos                rl.Vector2
	IsPlacedOnResource bool
	ExtractTime        float32
}

func (b *Building) Update() {
	b.ExtractTime -= float32(rl.GetFrameTime())

	if b.ExtractTime <= 0 {
		b.ExtractTime = 1.36
	}
}

func (b Building) Draw() {
	rl.DrawRectangle(int32(b.Pos.X), int32(b.Pos.Y), 50, 50, rl.Green)
}
