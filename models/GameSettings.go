package models

type GameSettings struct {
	Id      string   `json:"id"`
	Ruleset *Ruleset `json:"ruleset"`
	Map     string   `json:"map"`
	Timeout int      `json:"timeout"`
	Source  string   `json:"source"`
}
