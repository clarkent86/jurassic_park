package cage

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type responseObject struct {
	CageState  Cage
	StatusCode int
}

// Handler is responsible for defining a HTTP request route and corresponding handler.
type Handler struct {
	// Receives a route to modify, like adding path, methods, etc.
	Route func(r *mux.Route)

	// HTTP Handler
	Func http.HandlerFunc
}

// AddRoute adds the handler's route the to the router.
func (h Handler) AddRoute(r *mux.Router) {
	h.Route(r.NewRoute().HandlerFunc(h.Func))
}

func (park *Park) AddDinosaurToCageHandler(prefix string) Handler {
	return Handler{
		Route: func(r *mux.Route) {
			v1 := r.PathPrefix("/api/v1")
			v1.Path("add/dinosaur")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			cageName := r.URL.Query().Get("cageName")
			park.addDinosaurToCage(cageName, r.URL.Query().Get("dinosaurName"), r.URL.Query().Get("dinosaurSpecies"))

			json.NewEncoder(w).Encode(responseObject{
				CageState:  park.cages[cageName],
				StatusCode: 200,
			})
		},
	}
}

func (park *Park) NewCageHandler(w http.ResponseWriter, r *http.Request) Handler {
	return Handler{
		Route: func(r *mux.Route) {
			v1 := r.PathPrefix("/api/v1")
			v1.Path("add/cage")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			cageName := r.URL.Query().Get("cageName")
			// TODO: Deal with error
			capacity, _ := strconv.Atoi(r.URL.Query().Get("capacity"))
			park.addCage(cageName, capacity)

			json.NewEncoder(w).Encode(responseObject{
				CageState:  park.cages[cageName],
				StatusCode: 200,
			})
		},
	}

}

func (park *Park) PowerCageHandler(w http.ResponseWriter, r *http.Request) Handler {
	return Handler{
		Route: func(r *mux.Route) {
			v1 := r.PathPrefix("/api/v1")
			v1.Path("togglePower")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			cageName := r.URL.Query().Get("cageName")
			// TODO: Deal with error
			park.togglePower(cageName)

			json.NewEncoder(w).Encode(responseObject{
				CageState:  park.cages[cageName],
				StatusCode: 200,
			})
		},
	}
}
