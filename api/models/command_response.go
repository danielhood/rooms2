package models

type CommandResponse struct {
	User       string   `json:"user"`
	Command    string   `json:"command"`
	Subcommand string   `json:"subcommand"`
	Responses  []string `json:"responses"`
}
