package services

import (
	"github.com/danielhood/rooms2/api/models"
)

type CommandService interface {
	GetCommandResponse(command string, user string) (*models.CommandResponse, error)
}

func NewCommandService() CommandService {
	return &commandService{}
}

type commandService struct{}

func (c *commandService) GetCommandResponse(command string, user string) (*models.CommandResponse, error) {
	return &models.CommandResponse{
		User:       user,
		Command:    command,
		Subcommand: "test-test",
		Responses:  []string{"You are in a room.", "It is very dark."},
	}, nil
}
