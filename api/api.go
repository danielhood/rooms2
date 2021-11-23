package main

import (
	"log"
	"net/http"
	"os"

	"git.mills.io/prologic/bitcask"

	"github.com/danielhood/rooms2/api/handlers"
	"github.com/danielhood/rooms2/api/models"
	"github.com/danielhood/rooms2/api/repo"
	"github.com/danielhood/rooms2/api/security"
)

func generateDefaultUsers(userRepo repo.UserRepo) {
	if users, err := userRepo.GetAll(); err != nil {
		panic(err)
	} else {
		if len(users) == 0 {
			log.Print("Generating default users")
			// Initialize default users if no users currently exist
			userRepo.Add(&models.User{
				Username:  "admin",
				Password:  "admin",
				Roles:     []string{models.AdministratorRole},
				IsOnline:  false,
				IsEnabled: true,
			})

			userRepo.Add(&models.User{
				Username:  "test1",
				Password:  "test",
				Roles:     []string{models.UserRole},
				IsOnline:  false,
				IsEnabled: true,
			})

			userRepo.Add(&models.User{
				Username:  "test2",
				Password:  "test",
				Roles:     []string{models.UserRole},
				IsOnline:  false,
				IsEnabled: true,
			})

		} else {
			log.Print("Loaded users: ", len(users))
		}
	}
}

func createDefaultRoutes(
	userRepo repo.UserRepo,
) {

	pingHandler := handlers.NewPing()
	commandHandler := handlers.NewCommand(userRepo)
	tokenHandler := handlers.NewToken(userRepo)

	auth := security.NewAuthentication()

	http.Handle("/ping", pingHandler)
	http.Handle("/command", addMiddleware(commandHandler, auth.Authenticate))
	http.Handle("/token", tokenHandler)
}

func addMiddleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

func main() {
	log.Print("rooms2 api server starting")

	db, _ := bitcask.Open("./rooms2db")
	defer db.Close()

	storageManager := repo.NewStorageManager(db)

	userRepo := repo.NewUserRepo(storageManager)

	createDefaultRoutes(userRepo)
	generateDefaultUsers(userRepo)

	// openssl genrsa -out server.key 2048
	certPath := "server.pem"

	// openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
	keyPath := "server.key"

	if _, err := os.Stat(keyPath); err == nil {
		log.Print("Listening for connections on HTTPS port 8443")
		log.Fatal(http.ListenAndServeTLS(":8443", certPath, keyPath, nil))
	} else if os.IsNotExist(err) {
		log.Print("Listening for connections on HTTP port 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}

	log.Print("Terminating")
}
