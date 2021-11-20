module github.com/danielhood/rooms2/api/handlers

go 1.17

require (
	github.com/danielhood/rooms2/api/models v0.0.0-00010101000000-000000000000
	github.com/danielhood/rooms2/api/services v0.0.0-00010101000000-000000000000
)

require github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect

replace github.com/danielhood/rooms2/api/services => ../services

replace github.com/danielhood/rooms2/api/models => ../models
