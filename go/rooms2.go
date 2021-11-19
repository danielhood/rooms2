package main

import (
	"log"
	"net/http"
	"os"

	"rooms2-api/handlers"

	"git.mills.io/prologic/bitcask"
)

func createDefaultRoutes() {
	pingHandler := handlers.NewPing()

	http.Handle("/ping", pingHandler)
}

func main() {
	log.Print("rooms2 api server starting")

	createDefaultRoutes()

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

	db, _ := bitcask.Open("./rooms2db")
	defer db.Close()
	db.Put([]byte("Hello"), []byte("World"))
	val, _ := db.Get([]byte("Hello"))
	log.Print(string(val))

}
