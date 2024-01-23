package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CellType int

const (
	Empty CellType = iota
	Build
	Resource
)

type Cell struct {
	Type CellType
}

type Grid struct {
	cells map[int]map[int]Cell
}

const (
	screenWidth  = 1920
	screenHeight = 1080
)

var game = Game{
	messageQueue: MessageQueue{msgs: make([]*Message, 0)},
	resources:    make([]rl.Vector2, 0),
	buildings:    make([]Building, 0),
	debugMode:    false,
	camera: GameCamera{
		camera: rl.Camera2D{
			Zoom:     1.0,
			Rotation: 0.0,
			Offset:   rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2)),
		},
		trgt: snapToGrid(rl.NewVector2(0, 0)),
		zoom: 1.0,
	},
}

var grid = Grid{
	cells: make(map[int]map[int]Cell),
}

type Message struct {
	text string
	time float32
}

type MessageQueue struct {
	msgs []*Message
}

func main() {
	for i := -10000; i < 10000; i++ {
		grid.cells[i] = make(map[int]Cell)
	}

	rl.InitWindow(screenWidth, screenHeight, "Factory Game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	ps := PlacementSystem{}

	for i := 0; i < 100; i++ {
		randX := rl.GetRandomValue(-10000, 10000)
		randY := rl.GetRandomValue(-10000, 10000)

		randPoint := snapToGrid(rl.NewVector2(float32(randX), float32(randY)))

		dist := rl.Vector2Distance(rl.NewVector2(0, 0), randPoint)

		if dist < 300 {
			i--
			continue
		}

		game.resources = append(game.resources, randPoint)
		grid.cells[int(randPoint.X)][int(randPoint.Y)] = Cell{
			Type: Resource,
		}
	}

	newPoints := make([]rl.Vector2, 0)
	for _, point := range game.resources {
		x := point.X
		y := point.Y
		radius := float32(rl.GetRandomValue(50, 200))
		step := float32(50)
		for i := x - radius; i <= x+radius; i += step {
			for j := y - radius; j <= y+radius; j += step {
				dist := rl.Vector2Distance(rl.NewVector2(x, y), rl.NewVector2(i, j))
				if dist <= radius {
					newPoints = append(newPoints, snapToGrid(rl.NewVector2(i, j)))
					grid.cells[int(i)][int(j)] = Cell{
						Type: Resource,
					}
				}

			}
		}
	}

	game.resources = append(game.resources, newPoints...)

	for !rl.WindowShouldClose() {
		game.Update()

		rl.BeginDrawing()
		rl.BeginMode2D(game.camera.camera)

		for i := range game.buildings {
			game.buildings[i].Update()
		}

		rl.ClearBackground(rl.RayWhite)
		if game.debugMode {
			for i := -10000; i <= 10000; i += 50 {
				rl.DrawLine(int32(i), -10000, int32(i), 10000, rl.Black)
				rl.DrawLine(-10000, int32(i), 10000, int32(i), rl.Black)
			}
		}

		for i := 0; i < len(game.resources); i++ {
			rl.DrawRectangle(int32(game.resources[i].X), int32(game.resources[i].Y), 50, 50, rl.Blue)
		}

		ps.Update()

		for _, b := range game.buildings {
			b.Draw()
		}

		rl.DrawRectangleV(snapToGrid(rl.Vector2{X: 0, Y: 0}), snapToGrid(rl.Vector2{X: 100, Y: 100}), rl.Black)
		ps.Draw()

		rl.EndMode2D()

		game.DrawUI()

		rl.EndDrawing()
	}
}
