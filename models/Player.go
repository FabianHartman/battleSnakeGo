package models

type Player struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Latency        string          `json:"latency"`
	Health         int             `json:"health"`
	Head           *Coord          `json:"head"`
	Length         int             `json:"length"`
	Shout          string          `json:"shout"`
	Squad          string          `json:"squad"`
	Customizations *Customizations `json:"customizations"`
	Body           []*Coord        `json:"body"`
}
