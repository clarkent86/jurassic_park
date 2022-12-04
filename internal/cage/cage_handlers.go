package cage

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const v1 = "/api/v1"

type responseObject struct {
	CageState  Cage
	StatusCode int
}

type basicSuccess struct {
	Message    string
	StatusCode int
}

type errorResponse struct {
	Error      string
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
			v1 := r.PathPrefix(v1)
			v1.Path(prefix).Methods("POST")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			cageName := r.URL.Query().Get("cageName")
			err := park.addDinosaurToCage(cageName, r.URL.Query().Get("dinosaurName"), r.URL.Query().Get("dinosaurSpecies"))
			if err == nil {
				json.NewEncoder(w).Encode(responseObject{
					CageState:  park.cages[cageName],
					StatusCode: 200,
				})
			} else {
				json.NewEncoder(w).Encode(errorResponse{
					Error:      err.Error(),
					StatusCode: 400,
				})
			}
		},
	}
}

func (park *Park) RemoveDinosaurFromCageHandler(prefix string) Handler {
	return Handler{
		Route: func(r *mux.Route) {
			v1 := r.PathPrefix(v1)
			v1.Path(prefix).Methods("DELETE")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			cageName := r.URL.Query().Get("cageName")
			err := park.removeDinosaurFromCage(cageName, r.URL.Query().Get("dinosaurName"), r.URL.Query().Get("dinosaurSpecies"))
			if err == nil {
				json.NewEncoder(w).Encode(responseObject{
					CageState:  park.cages[cageName],
					StatusCode: 200,
				})
			} else {
				json.NewEncoder(w).Encode(errorResponse{
					Error:      err.Error(),
					StatusCode: 400,
				})
			}
		},
	}
}

func (park *Park) NewCageHandler(prefix string) Handler {
	return Handler{
		Route: func(r *mux.Route) {
			v1 := r.PathPrefix(v1)
			v1.Path(prefix).Methods("POST")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			cageName := r.URL.Query().Get("cageName")
			capacity, _ := strconv.Atoi(r.URL.Query().Get("capacity"))
			err := park.addCage(cageName, capacity)
			if err == nil {
				json.NewEncoder(w).Encode(responseObject{
					CageState:  park.cages[cageName],
					StatusCode: 200,
				})
			} else {
				json.NewEncoder(w).Encode(errorResponse{
					Error:      err.Error(),
					StatusCode: 400,
				})
			}
		},
	}
}

func (park *Park) RemoveCageHandler(prefix string) Handler {
	return Handler{
		Route: func(r *mux.Route) {
			v1 := r.PathPrefix(v1)
			v1.Path(prefix).Methods("DELETE")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			cageName := r.URL.Query().Get("cageName")
			err := park.removeCage(cageName)
			if err == nil {
				json.NewEncoder(w).Encode(basicSuccess{
					Message:    "cage successfully removed",
					StatusCode: 200,
				})
			} else {
				json.NewEncoder(w).Encode(errorResponse{
					Error:      err.Error(),
					StatusCode: 400,
				})
			}
		},
	}
}

func (park *Park) ToggleCageHandler(prefix string) Handler {
	return Handler{
		Route: func(r *mux.Route) {
			v1 := r.PathPrefix(v1)
			v1.Path(prefix).Methods("POST")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			cageName := r.URL.Query().Get("cageName")
			err := park.togglePower(cageName)
			if err == nil {
				json.NewEncoder(w).Encode(responseObject{
					CageState:  park.cages[cageName],
					StatusCode: 200,
				})
			} else {
				json.NewEncoder(w).Encode(errorResponse{
					Error:      err.Error(),
					StatusCode: 400,
				})
			}
		},
	}
}

func (park *Park) GetParkStatusHandler(prefix string) Handler {
	return Handler{
		Route: func(r *mux.Route) {
			v1 := r.PathPrefix(v1)
			v1.Path(prefix).Methods("GET")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(basicSuccess{
				Message:    park.String(),
				StatusCode: 200,
			})
		},
	}
}
