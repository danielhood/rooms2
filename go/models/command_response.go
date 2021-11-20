package models

// Quest defines a quest entity for our application
type CommandResponse struct {
	Command    string   `json:"command"`
	Subcommand string   `json:"subcommand"`
	Responses  []string `json:"responses"`
}
