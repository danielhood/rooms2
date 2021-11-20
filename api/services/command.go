package services

import (
	"github.com/danielhood/rooms2/api/models"
)

type CommandService interface {
	GetCommandResponse(command string) (*models.CommandResponse, error)
}

func NewCommandService() CommandService {
	return &commandService{}
}

type commandService struct{}

func (c *commandService) GetCommandResponse(command string) (*models.CommandResponse, error) {
	return &models.CommandResponse{
		Command:    "test",
		Subcommand: "test-test",
		Responses:  []string{"You are in a room.", "It is very dark."},
	}, nil
}
