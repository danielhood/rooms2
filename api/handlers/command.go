package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/danielhood/rooms2/api/services"
)

// Command holds handler structure
type Command struct {
	svc services.CommandService
}

// NewCommand creates an instance of Command
func NewCommand() *Command {
	return &Command{services.NewCommandService()}
}

func (h *Command) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.enableCors(&w)

	switch req.Method {
	case "OPTIONS":

		log.Print("/token:OPTIONS")

		if req.Header.Get("Access-Control-Request-Method") != "" {
			w.Header().Set("Allow", req.Header.Get("Access-Control-Request-Method"))
			w.Header().Set("Access-Control-Allow-Methods", req.Header.Get("Access-Control-Request-Method"))
		}

		w.Header().Set("Access-Control-Allow-Headers", "authorization,access-control-allow-origin,content-type")

	case "GET":
		log.Print("/command:GET")
		log.Print("GET params were:", req.URL.Query())

		user := req.URL.Query().Get("u")
		commandText := req.URL.Query().Get("c")

		if len(commandText) == 0 || len(user) == 0 {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		} else {
			commandResponse, err := h.svc.GetCommandResponse(commandText, user)

			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			responseBytes, _ := json.Marshal(commandResponse)
			w.Write(responseBytes)
		}

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Command) enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
