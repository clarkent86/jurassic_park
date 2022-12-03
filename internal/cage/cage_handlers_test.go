package cage

import (
	"net/http/httptest"
	"testing"

	"github.com/clarkent86/jurassic_park/internal/dinosaur"
	"github.com/gorilla/mux"
)

func TestAddDinosaurToCageHandler(t *testing.T) {
	tests := []struct {
		description   string
		dinosaurToAdd dinosaur.Dinosaur
		expectedCage  Cage
		expectedError error
		initialCage   Cage
	}{{ // successful add carnivorous dinosaur to empty powered cage
		description:   testAddDescriptions[0],
		dinosaurToAdd: testFirstCarnivorousDinosaur,
		expectedCage: Cage{
			Capacity: 1,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[0],
			PowerStatus: "ACTIVE",
			Type:        "carnivorous",
		},
		expectedError: nil,
		initialCage: Cage{
			Capacity:    1,
			Dinosaurs:   []dinosaur.Dinosaur{},
			Name:        testAddDescriptions[0],
			PowerStatus: "ACTIVE",
			Type:        "",
		},
	}, { // successful add herbivorous dinosaur to empty powered cage
		description:   testAddDescriptions[1],
		dinosaurToAdd: testFirstHerbivorousDinosaur,
		expectedCage: Cage{
			Capacity: 1,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstHerbivorousDinosaur,
			},
			Name:        testAddDescriptions[1],
			PowerStatus: "ACTIVE",
			Type:        "herbivorous",
		},
		expectedError: nil,
		initialCage: Cage{
			Capacity:    1,
			Dinosaurs:   []dinosaur.Dinosaur{},
			Name:        testAddDescriptions[1],
			PowerStatus: "ACTIVE",
			Type:        "",
		},
	}}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := mux.NewRouter()

			// AddDinosaurToCage("/add/dinosaur").AddRoute(r)

			r.ServeHTTP(w, httptest.NewRequest("GET", "/api/v1/add/dinosaur", nil))

			// if w.Code != http.StatusOK {

			// }

			// err := test.initialCage.AddDinosaurToCageHandler(test.dinosaurToAdd)
			// assert.Equal(t, test.expectedCage, test.initialCage)
			// if assert.Error(t, err) {
			// 	assert.Equal(t, test.expectedError, err)
			// }
		})
	}
}
