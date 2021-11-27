package services

import (
	"github.com/danielhood/rooms2/api/models"
	"github.com/danielhood/rooms2/api/repo"
)

func ProcessCommand(command *models.Command, userRepo repo.UserRepo) (*models.CommandResponse, error) {
	switch command.CommandType {
	case models.CommandType_Help:
		return helpResponse(command.User, isAdminUser(command.User, userRepo))
	case models.CommandType_Enter:
		return enterResponse(command.User)
	case models.CommandType_Exit:
		return exitResponse(command.User)
	case models.CommandType_User:
		if !isAdminUser(command.User, userRepo) {
			return unknownResponse(command.User)
		}

		switch command.Target {
		case "ADD":
			return addUser(command, userRepo)
		case "PASSWORD":
			return changePassword(command, userRepo)
		}
	}

	return unknownResponse(command.User)
}

func helpResponse(user string, isAdmin bool) (*models.CommandResponse, error) {
	responseStrings := []string{}

	if isAdmin {
		responseStrings = append(responseStrings,
			"user <add | password> username password",
		)
	}

	responseStrings = append(responseStrings,
		"enter, exit, n, north, s, south",
		"look <direction>",
		"look at <object>",
		"take <item>",
		"drop <item>",
	)
	return &models.CommandResponse{
		User:      user,
		Command:   models.CommandType_Help,
		Responses: responseStrings,
	}, nil
}

func unknownResponse(user string) (*models.CommandResponse, error) {
	return &models.CommandResponse{
		User:      user,
		Command:   models.CommandType_None,
		Responses: []string{"You look around somewhat confused."},
	}, nil
}

func exitResponse(user string) (*models.CommandResponse, error) {
	return &models.CommandResponse{
		User:      user,
		Command:   models.CommandType_Exit,
		Responses: []string{"You have left the world and are floating in a void."},
	}, nil
}

func enterResponse(user string) (*models.CommandResponse, error) {
	return &models.CommandResponse{
		User:      user,
		Command:   models.CommandType_Enter,
		Responses: []string{"You have left the void and are standing in a room."},
	}, nil
}

func isAdminUser(username string, userRepo repo.UserRepo) bool {
	user, _ := userRepo.GetByUsername(username)

	return user.HasRole(models.AdministratorRole)
}

func addUser(command *models.Command, userRepo repo.UserRepo) (*models.CommandResponse, error) {
	_ = userRepo.Add(&models.User{
		Username:  command.Param1,
		Password:  command.Param2,
		Roles:     []string{models.UserRole},
		IsOnline:  false,
		IsEnabled: true,
	})

	return &models.CommandResponse{
		User:      command.User,
		Command:   models.CommandType_User,
		Responses: []string{"User added."},
	}, nil
}

func changePassword(command *models.Command, userRepo repo.UserRepo) (*models.CommandResponse, error) {
	user, _ := userRepo.GetByUsername(command.Param1)

	user.Password = command.Param2

	_ = userRepo.Add(user)

	return &models.CommandResponse{
		User:      command.User,
		Command:   models.CommandType_User,
		Responses: []string{"Password changed."},
	}, nil
}
