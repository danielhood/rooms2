module github.com/danielhood/rooms2/api/services

go 1.17

require (
	github.com/danielhood/rooms2/api/models v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.3.0
)

replace github.com/danielhood/rooms2/api/models => ../models
