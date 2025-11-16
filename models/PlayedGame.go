package models

type Game struct {
	ID       string  `json:"id"`
	Map      string  `json:"map"`
	Finished bool    `json:"finished"`
	Height   int     `json:"height"`
	Width    int     `json:"width"`
	Score    int     `json:"score"`
	Moves    []Coord `json:"moves"`
}

var AllGames []Game

func NewGame(id string, gameMap string, height int, width int) *Game {
	game := Game{
		ID:     id,
		Map:    gameMap,
		Height: height,
		Width:  width,
	}

	AllGames = append(AllGames, game)

	return &game
}

func SearchGameByID(id string) *Game {
	for i := range AllGames {
		if AllGames[i].ID == id {
			return &AllGames[i]
		}
	}

	return nil
}

func RemoveGameByID(id string) bool {
	index := -1

	for i := range AllGames {
		if AllGames[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		return false
	}

	AllGames = append(AllGames[:index], AllGames[index+1:]...)

	return true
}
