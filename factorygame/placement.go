package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlacementSystem struct{
	ghost rl.Vector2
}

func (ps *PlacementSystem) Update() {
	ps.ghost = snapToGrid(getWorldMousePos(game.GetCamera()))

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		ps.handlePlacement(ps.ghost)
	}
}

func (ps *PlacementSystem) Draw() {
	rl.DrawRectangleV(ps.ghost, rl.NewVector2(50, 50), rl.Red)
}

func (ps *PlacementSystem) handlePlacement(ghost rl.Vector2) {
	hasCollision := false
	placedOnResource := false

	building := Building{Pos: ghost, IsPlacedOnResource: false}

	homePos := snapToGrid(rl.NewVector2(0, 0))

	homeRect := rl.NewRectangle(homePos.X, homePos.Y, 100, 100)
	placementRec := rl.NewRectangle(ghost.X, ghost.Y, 50, 50)

	if rl.CheckCollisionRecs(homeRect, placementRec) {
		hasCollision = true
		game.messageQueue.msgs = append(game.messageQueue.msgs, &Message{text: "Collision!", time: 3})
	}

	if !hasCollision {
		for _, b := range game.buildings {
			rect1 := rl.NewRectangle(b.Pos.X, b.Pos.Y, 50, 50)

			if rl.CheckCollisionRecs(rect1, placementRec) {
				hasCollision = true
				game.messageQueue.msgs = append(game.messageQueue.msgs, &Message{text: "Collision!", time: 3})
			}
		}
	}

	if !hasCollision {
		for _, b := range game.resources {
			rect1 := rl.NewRectangle(b.X, b.Y, 50, 50)

			if rl.CheckCollisionRecs(rect1, placementRec) {
				placedOnResource = true
			}
		}

		building.IsPlacedOnResource = placedOnResource

		if placedOnResource {
			game.buildings = append(game.buildings, building)
			game.messageQueue.msgs = append(game.messageQueue.msgs, &Message{text: "Placed on resource!", time: 3})
		} else {
			game.buildings = append(game.buildings, building)
			game.messageQueue.msgs = append(game.messageQueue.msgs, &Message{text: "Spawned!", time: 3})
		}

		grid.cells[int(ghost.X)][int(ghost.Y)] = Cell{
			Type: Build,
		}

		cellsForHeight := (building.Pos.Y)/50
		cellsForWidth := (building.Pos.X)/50

		fmt.Println(building.Pos)
		fmt.Println(cellsForHeight, cellsForWidth)
	}
}
