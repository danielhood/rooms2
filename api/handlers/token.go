package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/danielhood/rooms2/api/models"
	"github.com/danielhood/rooms2/api/repo"
	"github.com/danielhood/rooms2/api/services"
)

// Token contains strucutre of a token handler
type Token struct {
	svc      services.TokenService
	userRepo repo.UserRepo
}

// NewToken creates new handler for tokens
func NewToken(
	userRepo repo.UserRepo,
) *Token {
	return &Token{
		svc:      services.NewTokenService(userRepo),
		userRepo: userRepo,
	}
}

// Handler will return tokens
func (t *Token) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t.enableCors(&w)

	log.Print("/token:", req.Method)

	switch req.Method {
	case "OPTIONS":
		w.Header().Set("Allow", "GET,POST")
		w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin,content-type")

	case "GET":
		log.Print("GET params were:", req.URL.Query())

		username := req.URL.Query().Get("u")
		password := req.URL.Query().Get("p")

		if len(username) == 0 || len(password) == 0 {
			http.Error(w, "Invalid query params", http.StatusBadRequest)
			return
		}

		t.generateTokenResponse(w, username, password)

	case "POST":
		requestBody, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			http.Error(w, "Unable to parse request body", http.StatusBadRequest)
			return
		}

		if len(requestBody) == 0 {
			http.Error(w, "Empty TokenRequest passed", http.StatusBadRequest)
			return
		}

		var tokenRequest models.TokenRequest
		if err = json.Unmarshal(requestBody, &tokenRequest); err != nil {
			http.Error(w, "Unable to parse token request json", http.StatusBadRequest)
			return
		}

		if len(tokenRequest.Username) == 0 || len(tokenRequest.Password) == 0 {
			http.Error(w, "Invalid request body: missing requierd keys", http.StatusBadRequest)
			return
		}

		t.generateTokenResponse(w, tokenRequest.Username, tokenRequest.Password)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (t *Token) generateTokenResponse(w http.ResponseWriter, username string, password string) {

	token, err := t.svc.ProcessUserLogin(username, password)

	if err != nil {
		http.Error(w, "Failed to verify user credentials", http.StatusUnauthorized)
		return
	}

	var tokenResponse models.TokenResponse
	tokenResponse.Token = token

	var tokenBytes []byte
	tokenBytes, _ = json.Marshal(tokenResponse)

	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenBytes)
}

func (t *Token) enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
