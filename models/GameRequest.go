package models

import (
	"math"
	"slices"

	"battleSnakeGo/helpers"
)

type GameRequest struct {
	Game  *GameSettings `json:"game"`
	Turn  int           `json:"turn"`
	Board *Board        `json:"board"`
	You   *Player       `json:"you"`
}

func (this *GameRequest) GenerateMove() string {
	head := this.You.Head
	bodyParts := this.You.Body
	board := this.Board

	possibilities := []string{"down", "up", "left", "right"}

	if head.X < bodyParts[1].X {
		possibilities = helpers.RemoveStringFromSlice(possibilities, "right")
	} else if head.X == board.Width-1 {
		possibilities = helpers.RemoveStringFromSlice(possibilities, "right")
	}

	if head.X > bodyParts[1].X {
		possibilities = helpers.RemoveStringFromSlice(possibilities, "left")
	} else if head.X == 0 {
		possibilities = helpers.RemoveStringFromSlice(possibilities, "left")
	}

	if head.Y > bodyParts[1].Y {
		possibilities = helpers.RemoveStringFromSlice(possibilities, "down")
	} else if head.Y == 0 {
		possibilities = helpers.RemoveStringFromSlice(possibilities, "down")
	}

	if head.Y < bodyParts[1].Y {
		possibilities = helpers.RemoveStringFromSlice(possibilities, "up")
	} else if head.Y == board.Height-1 {
		possibilities = helpers.RemoveStringFromSlice(possibilities, "up")
	}

	for _, bodyPart := range bodyParts[1:] {
		_, possibilities = removeImpossibleMoves(head, bodyPart, possibilities)
	}

	var removedPossibilitiesEnemy []string

	for _, enemy := range board.Snakes {
		if enemy.ID == this.You.ID {
			continue
		}

		for _, bodyPartEnemy := range enemy.Body {
			removedMove, updatedPossibilities := removeImpossibleMoves(head, bodyPartEnemy, possibilities)
			possibilities = updatedPossibilities

			if removedMove != "" {
				removedPossibilitiesEnemy = append(removedPossibilitiesEnemy, removedMove)
			}

			removedMove, possibilities = removeImpossibleMoves(head, &Coord{bodyPartEnemy.X - 1, bodyPartEnemy.Y}, possibilities)
			if removedMove != "" {
				removedPossibilitiesEnemy = append(removedPossibilitiesEnemy, removedMove)
			}

			removedMove, possibilities = removeImpossibleMoves(head, &Coord{bodyPartEnemy.X + 1, bodyPartEnemy.Y}, possibilities)
			if removedMove != "" {
				removedPossibilitiesEnemy = append(removedPossibilitiesEnemy, removedMove)
			}

			removedMove, possibilities = removeImpossibleMoves(head, &Coord{bodyPartEnemy.X, bodyPartEnemy.Y + 1}, possibilities)
			if removedMove != "" {
				removedPossibilitiesEnemy = append(removedPossibilitiesEnemy, removedMove)
			}

			removedMove, possibilities = removeImpossibleMoves(head, &Coord{bodyPartEnemy.X, bodyPartEnemy.Y - 1}, possibilities)
			if removedMove != "" {
				removedPossibilitiesEnemy = append(removedPossibilitiesEnemy, removedMove)
			}
		}
	}

	var closestFood Coord
	minDistance := math.MaxInt

	for _, food := range board.Food {
		distance := helpers.AbsValue(food.X-head.X) + helpers.AbsValue(food.Y-head.Y)
		if distance < minDistance {
			minDistance = distance
			closestFood = food
		}
	}

	var move string
	if closestFood != (Coord{}) {
		if closestFood.X < head.X && slices.Contains(possibilities, "left") {
			move = "left"
		} else if closestFood.X > head.X && slices.Contains(possibilities, "right") {
			move = "right"
		} else if closestFood.Y < head.Y && slices.Contains(possibilities, "down") {
			move = "down"
		} else if closestFood.Y > head.Y && slices.Contains(possibilities, "up") {
			move = "up"
		} else if len(possibilities) > 0 {
			move = possibilities[0]
		}
	} else {
		if len(possibilities) > 0 {
			move = possibilities[0]
		} else if len(removedPossibilitiesEnemy) > 0 {
			move = removedPossibilitiesEnemy[0]
		} else {
			move = "up"
		}
	}

	return move
}

func removeImpossibleMoves(head *Coord, bodyPart *Coord, possibilities []string) (string, []string) {
	var removedMove string
	for i, move := range possibilities {
		switch {
		case head.X == bodyPart.X-1 && head.Y == bodyPart.Y && move == "right":
			removedMove = "right"
			possibilities = append(possibilities[:i], possibilities[i+1:]...)
		case head.X == bodyPart.X+1 && head.Y == bodyPart.Y && move == "left":
			removedMove = "left"
			possibilities = append(possibilities[:i], possibilities[i+1:]...)
		case head.Y == bodyPart.Y+1 && head.X == bodyPart.X && move == "down":
			removedMove = "down"
			possibilities = append(possibilities[:i], possibilities[i+1:]...)
		case head.Y == bodyPart.Y-1 && head.X == bodyPart.X && move == "up":
			removedMove = "up"
			possibilities = append(possibilities[:i], possibilities[i+1:]...)
		}
	}
	return removedMove, possibilities
}
