package services

import (
	"strings"

	"github.com/danielhood/rooms2/api/models"
)

func ParseCommand(commandText string) *models.Command {
	commandTokens := strings.Fields(strings.ToUpper(commandText))
	commandTokensOrig := strings.Fields(commandText)

	if len(commandTokens) == 0 {
		return invalidCommand()
	}

	switch commandTokens[0] {
	case models.CommandType_Exit:
		return simpleCommand(commandTokens[0])
	case models.CommandType_Enter:
		return simpleCommand(commandTokens[0])
	case models.CommandType_Help:
		return simpleCommand(commandTokens[0])
	case models.CommandType_Look:
		if len(commandTokens) > 1 {
			return lookCommand("")
		} else {
			if commandTokens[1] == "AT" {
				return lookAtCommand(commandTokens[2])
			} else {
				return lookCommand(commandTokens[1])
			}
		}
	case models.CommandType_Take:
		return takeCommand(commandTokens)
	case models.CommandType_Drop:
		return dropCommand(commandTokens)
	case models.CommandType_User:
		return userCommand(commandTokens, commandTokensOrig)
	}

	return invalidCommand()
}

func parseDirection(directionText string) string {
	switch strings.ToUpper(directionText) {
	case models.Dir_N, "NORTH":
		return models.Dir_N
	case models.Dir_NE, "NORTHEAST":
		return models.Dir_N
	case models.Dir_E, "EAST":
		return models.Dir_N
	case models.Dir_SE, "SOUTHEAST":
		return models.Dir_N
	case models.Dir_S, "SOUTH":
		return models.Dir_N
	case models.Dir_SW, "SOUTHWEST":
		return models.Dir_N
	case models.Dir_W, "WEST":
		return models.Dir_N
	case models.Dir_NW, "NORTHWEST":
		return models.Dir_N
	case models.Dir_U, "UP":
		return models.Dir_N
	case models.Dir_D, "DOWN":
		return models.Dir_N
	default:
		return ""
	}
}

func invalidCommand() *models.Command {
	return &models.Command{
		CommandType: models.CommandType_None,
	}
}

func simpleCommand(commandType string) *models.Command {
	return &models.Command{
		CommandType: commandType,
	}
}

func lookAtCommand(target string) *models.Command {
	return &models.Command{
		CommandType: models.CommandType_LookAt,
		Target:      target,
	}
}

func lookCommand(directionText string) *models.Command {
	dir := parseDirection((directionText))

	if dir == "" {
		return &models.Command{
			CommandType: models.CommandType_Look,
		}
	} else {
		return &models.Command{
			CommandType: models.CommandType_Look,
			Direction:   dir,
		}
	}
}

func takeCommand(commandTokens []string) *models.Command {
	if len(commandTokens) < 2 {
		return invalidCommand()
	}

	return &models.Command{
		CommandType: models.CommandType_Take,
		Target:      commandTokens[1],
	}
}

func dropCommand(commandTokens []string) *models.Command {
	if len(commandTokens) < 2 {
		return invalidCommand()
	}

	return &models.Command{
		CommandType: models.CommandType_Drop,
		Target:      commandTokens[1],
	}
}

func userCommand(commandTokens []string, commandTokensOrig []string) *models.Command {
	if len(commandTokens) < 3 {
		return invalidCommand()
	}

	switch commandTokens[1] {
	case "ADD":
		if len(commandTokens) < 4 {
			return invalidCommand()
		}

		return &models.Command{
			CommandType: models.CommandType_User,
			Target:      "ADD",
			Param1:      commandTokensOrig[2], // Username
			Param2:      commandTokensOrig[3], // Password
		}
	case "PASSWORD":
		if len(commandTokens) < 4 {
			return invalidCommand()
		}

		return &models.Command{
			CommandType: models.CommandType_User,
			Target:      "PASSWORD",
			Param1:      commandTokensOrig[2], // Username
			Param2:      commandTokensOrig[3], // Password
		}
	}

	return invalidCommand()
}
