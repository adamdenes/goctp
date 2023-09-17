package server

import (
	"log"
	"net/http"
	"text/template"

	"github.com/asynkron/protoactor-go/actor"
)

type Server struct {
	ctx     actor.Context
	address string
}

func New(address string) actor.Producer {
	return func() actor.Actor {
		return &Server{
			address: address,
		}
	}
}

func (a *Server) Receive(c actor.Context) {
	switch msg := c.Message().(type) {
	case *actor.Started:
		_ = msg
		a.ctx = c
		a.start()
	case *actor.Stopped:
		// stop DB
		// stop exchange stream
		// stop http server
	}
}

// Define a handler function for rendering the index page
func (a *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("www/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template with data (if needed)
	data := struct {
		Title string
		Body  string
	}{
		Title: "goctp",
		Body:  "My Crpyto Trading Platform",
	}

	// Render the template
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (a *Server) start() {
	http.HandleFunc("/", a.indexHandler)

	log.Printf("Server is running on %s...\n", a.address)
	err := http.ListenAndServe(a.address, nil)
	if err != nil {
		log.Printf("Server error: %v\n", err)
	}
}
