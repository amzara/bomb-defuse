package server

import (
	"fmt"
	"log"
	"net/http"
)

type Application struct {
	Port string
}

func New() *Application {
	return &Application{Port: ":1234"}
} //maybe better to change port number in func

func (app *Application) StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthCheck", healthCheck)

	go func() {
		fmt.Printf("Starting server at port %s", app.Port)
		err := http.ListenAndServe(app.Port, mux)
		if err != nil {
			log.Fatal(err)
		}
	}() //anonymous functions need this
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, "Server is up and running")

}
