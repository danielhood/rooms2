module github.com/danielhood/rooms2/api

go 1.17

require (
	git.mills.io/prologic/bitcask v1.0.2
	github.com/danielhood/rooms2/api/handlers v0.0.0-00010101000000-000000000000
)

require (
	github.com/abcum/lcp v0.0.0-20201209214815-7a3f3840be81 // indirect
	github.com/danielhood/rooms2/api/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/danielhood/rooms2/api/services v0.0.0-00010101000000-000000000000 // indirect
	github.com/gofrs/flock v0.8.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/plar/go-adaptive-radix-tree v1.0.4 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/exp v0.0.0-20200228211341-fcea875c7e85 // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/danielhood/rooms2/api/models => ./models

replace github.com/danielhood/rooms2/api/services => ./services

replace github.com/danielhood/rooms2/api/handlers => ./handlers
