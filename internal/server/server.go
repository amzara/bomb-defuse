package server

import (
	"bomb-defuse/internal/services/bomb"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Application struct {
	Port string
	Bomb *bomb.Bomb
}

func New() *Application {
	return &Application{Port: ":1234"}
} //maybe better to change port number in func

func (app *Application) StartServer(b *bomb.Bomb) {
	mux := http.NewServeMux()
	mux.HandleFunc("/plant", func(w http.ResponseWriter, r *http.Request) {
		b.Plant(10 * time.Second)
	})

	mux.HandleFunc("/defuse", func(w http.ResponseWriter, r *http.Request) {
		b.Defuse()
	})

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
