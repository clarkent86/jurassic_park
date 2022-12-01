package cage

import (
	"errors"
	"testing"

	"github.com/clarkent86/jurassic_park/internal/dinosaur"
	"github.com/stretchr/testify/assert"
)

var testAddDescriptions = []string{
	"successful add carnivorous dinosaur to empty powered cage",
	"successful add herbivorous dinosaur to empty powered cage",
	"failure to add dinosaur to down cage",
	"successful add to cage with existing dinosaur",
	"failure to add to cage at max capacity",
	"successful add like carnivorous dinosaur to existing carnivorous cage",
	"failure add carnivorous dinosaur to existing carnivorous cage with unlike species",
	"successful add herbivorous dinosaur to existing herbivorous cage with like species",
	"successful add herbivorous dinosaur to existing herbivorous cage with unlike species",
}

var testFirstCarnivorousDinosaur = dinosaur.Dinosaur{Name: "T-Rex", Species: "Tyrannosaurus"}
var testSecondCarnivorousDinosaur = dinosaur.Dinosaur{Name: "V-Raptor", Species: "Velociraptor"}
var testHerbivorousDinosaur = dinosaur.Dinosaur{Name: "B-Saurus", Species: "Brachiosaurus"}

func TestInitDinosaur(t *testing.T) {
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
		dinosaurToAdd: testHerbivorousDinosaur,
		expectedCage: Cage{
			Capacity: 1,
			Dinosaurs: []dinosaur.Dinosaur{
				testHerbivorousDinosaur,
			},
			Name:        testAddDescriptions[1],
			PowerStatus: "ACTIVE",
			Type:        "herbivorous",
		},
		expectedError: errors.New("[internal][cage] cannot add a dinosaur to a down cage"),
		initialCage: Cage{
			Capacity:    1,
			Dinosaurs:   []dinosaur.Dinosaur{},
			Name:        testAddDescriptions[1],
			PowerStatus: "ACTIVE",
			Type:        "",
		},
	}, { // failure to add dinosaur to down cage
		description:   testAddDescriptions[2],
		dinosaurToAdd: testFirstCarnivorousDinosaur,
		expectedCage: Cage{
			Capacity:    1,
			Dinosaurs:   []dinosaur.Dinosaur{},
			Name:        testAddDescriptions[2],
			PowerStatus: "DOWN",
			Type:        "",
		},
		expectedError: errors.New("[internal][cage] cannot add a dinosaur to a down cage"),
		initialCage: Cage{
			Capacity:    1,
			Dinosaurs:   []dinosaur.Dinosaur{},
			Name:        testAddDescriptions[2],
			PowerStatus: "DOWN",
			Type:        "",
		},
	}, { // successful add to cage with existing dinosaur
		description:   testAddDescriptions[3],
		dinosaurToAdd: dinosaur.Dinosaur{Name: "T-Rex", Species: "tyrannosaurus"},
		expectedCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[3],
			PowerStatus: "ACTIVE",
			Type:        "carnivorous",
		},
		expectedError: nil,
		initialCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[3],
			PowerStatus: "ACTIVE",
			Type:        "carnivorous",
		},
	}, { // failure to add to cage at max capacity
		description:   testAddDescriptions[4],
		dinosaurToAdd: dinosaur.Dinosaur{Name: "T-Rex", Species: "tyrannosaurus"},
		expectedCage: Cage{
			Capacity: 1,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[4],
			PowerStatus: "ACTIVE",
			Type:        "carnivorous",
		},
		expectedError: errors.New("[internal][cage] cannot add dinosaur - cage is at max capacity"),
		initialCage: Cage{
			Capacity: 1,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[4],
			PowerStatus: "ACTIVE",
			Type:        "carnivorous",
		},
	}}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.initialCage.AddDinosaurToCage(test.dinosaurToAdd)
			assert.Equal(t, test.expectedCage, test.initialCage)
			if assert.Error(t, err) {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}
