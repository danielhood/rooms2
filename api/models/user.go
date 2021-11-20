package models

const (
	// AdministratorRole defines role of administrator
	AdministratorRole = "administrator"
	UserRole          = "user"
)

// User defines a user for our application
type User struct {
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Roles     []string `json:"roles"`
	IsOnline  bool     `json:"isonline"`
	IsEnabled bool     `json:"isenabled"`
}

// HasRole returns true if the user is in the role
func (u *User) HasRole(roleName string) bool {
	for _, role := range u.Roles {
		if role == roleName {
			return true
		}
	}
	return false
}
