package models

type Snake struct {
	Customizations *Customizations `json:"customizations"`
	APIVersion     string          `json:"apiversion"`
	Author         string          `json:"author"`
	Name           string          `json:"name"`
}

var CurrentSnake = Snake{
	Customizations: &Customizations{
		Color: "#888888",
		Head:  "default",
		Tail:  "default",
	},
	APIVersion: "1",
	Author:     "Fabian Hartman",
	Name:       "epic snake",
}
