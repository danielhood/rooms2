package models

const (
	CommandType_None = "NONE"

	// User commands
	CommandType_Help   = "HELP"
	CommandType_Exit   = "EXIT"
	CommandType_Enter  = "ENTER"
	CommandType_Move   = "MOVE"
	CommandType_Look   = "LOOK"
	CommandType_LookAt = "LOOKAT"
	CommandType_Take   = "TAKE"
	CommandType_Drop   = "DROP"

	// Admin commands
	CommandType_User = "USER"
)

const (
	Dir_N  = "N"
	Dir_NE = "NE"
	Dir_E  = "E"
	Dir_SE = "SE"
	Dir_S  = "S"
	Dir_SW = "SW"
	Dir_W  = "W"
	Dir_NW = "NW"
	Dir_U  = "U"
	Dir_D  = "D"
)

type Command struct {
	User        string
	CommandType string
	Direction   string
	Target      string
	Param1      string
	Param2      string
}
