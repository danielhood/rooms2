module github.com/danielhood/rooms2/go/handlers

go 1.17

replace github.com/danielhood/rooms2/go/services => ../services

replace github.com/danielhood/rooms2/go/models => ../models

require github.com/danielhood/rooms2/go/services v0.0.0-00010101000000-000000000000

require github.com/danielhood/rooms2/go/models v0.0.0-00010101000000-000000000000 // indirect
