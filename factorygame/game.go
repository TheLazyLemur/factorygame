package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameCamera struct {
	camera rl.Camera2D
	trgt   rl.Vector2
	zoom   float32
	hasRun bool
}

func (c *GameCamera) Update() {
	c.camera.Target = c.trgt
	c.camera.Zoom = c.zoom
	if !c.hasRun {
		c.hasRun = true
		c.camera.Offset = rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2))
	}
}

type Game struct {
	camera       GameCamera
	messageQueue MessageQueue
	resources    []rl.Vector2
	buildings    []Building
	debugMode    bool
}

func (g *Game) GetCamera() rl.Camera2D {
	return g.camera.camera
}

func (g *Game) Update() {
	for i := 0; i < len(g.messageQueue.msgs); i++ {
		g.messageQueue.msgs[i].time -= 1 * rl.GetFrameTime()
		if g.messageQueue.msgs[i].time <= 0 {
			g.messageQueue.msgs = append(g.messageQueue.msgs[:i], g.messageQueue.msgs[i+1:]...)
		}
	}

	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		g.camera.trgt.X += 200 * float32(rl.GetFrameTime())
	}

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		g.camera.trgt.X -= 200 * float32(rl.GetFrameTime())
	}

	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		g.camera.trgt.Y -= 200 * float32(rl.GetFrameTime())
	}

	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		g.camera.trgt.Y += 200 * float32(rl.GetFrameTime())
	}

	scroll := rl.GetMouseWheelMove()

	if scroll > 0 {
		g.camera.zoom += 0.01 * 100 * 3 * float32(rl.GetFrameTime())
		if g.camera.zoom > 10.0 {
			g.camera.zoom = 10.0
		}
	}

	if scroll < 0 {
		g.camera.zoom -= 0.01 * 100 * 3 * float32(rl.GetFrameTime())
		if g.camera.zoom < 0.01 {
			g.camera.zoom = 0.01
		}
	}

	if rl.IsKeyPressed(rl.KeyF1) {
		g.debugMode = !g.debugMode
	}

	g.camera.Update()
}

func (g *Game) DrawUI() {
	if g.debugMode {
		for i := 0; i < len(g.messageQueue.msgs); i++ {
			rl.DrawText(g.messageQueue.msgs[i].text, 10, 30+(int32(i)*20), 20, rl.Black)
		}

		rl.DrawFPS(10, 10)
		rl.DrawText(fmt.Sprintf("Zoom: %.2f", g.camera.camera.Zoom), 10, 30, 20, rl.Black)
	}
}
