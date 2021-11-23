package services

import (
	"github.com/danielhood/rooms2/api/models"
	"github.com/danielhood/rooms2/api/repo"
)

type CommandService interface {
	GetCommandResponse(command string, user string) (*models.CommandResponse, error)
}

func NewCommandService(
	userRepo repo.UserRepo,
) CommandService {
	return &commandService{
		userRepo: userRepo,
	}
}

type commandService struct {
	userRepo repo.UserRepo
}

func (c *commandService) GetCommandResponse(commandText string, user string) (*models.CommandResponse, error) {

	command := ParseCommand(commandText)

	command.User = user

	return ProcessCommand(command, c.userRepo)
}
