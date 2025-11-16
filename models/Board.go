package models

type Board struct {
	Height  int      `json:"height"`
	Width   int      `json:"width"`
	Food    []Coord  `json:"food"`
	Hazards []Coord  `json:"hazards"`
	Snakes  []Player `json:"snakes"`
}
