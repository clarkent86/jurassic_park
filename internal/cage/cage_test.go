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
	"successful add like carnivorous dinosaur to existing carnivorous cage",
	"failure to add to cage at max capacity",
	"failure add carnivorous dinosaur to existing carnivorous cage with unlike species",
	"successful add herbivorous dinosaur to existing herbivorous cage with like species",
	"successful add herbivorous dinosaur to existing herbivorous cage with unlike species",
	"failure to add carnivorous dinosaur to existing herbivorous cage",
	"failure to add herbivorous dinosaur to existing carnivorous cage",
}

var testFirstCarnivorousDinosaur = dinosaur.Dinosaur{Name: "T-Rex", Species: "Tyrannosaurus"}
var testSecondCarnivorousDinosaur = dinosaur.Dinosaur{Name: "V-Raptor", Species: "Velociraptor"}
var testFirstHerbivorousDinosaur = dinosaur.Dinosaur{Name: "B-Saurus", Species: "Brachiosaurus"}
var testSecondHerbivorousDinosaur = dinosaur.Dinosaur{Name: "S-Saurus", Species: "Stegosaurus"}

func TestAddDinosaurToCage(t *testing.T) {
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
			Type:        "Carnivorous",
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
			Type:        "Herbivorous",
		},
		expectedError: nil,
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
		expectedError: errors.New(errPrefix + errDownCage),
		initialCage: Cage{
			Capacity:    1,
			Dinosaurs:   []dinosaur.Dinosaur{},
			Name:        testAddDescriptions[2],
			PowerStatus: "DOWN",
			Type:        "",
		},
	}, { // successful add to cage with existing dinosaur
		description:   testAddDescriptions[3],
		dinosaurToAdd: testFirstCarnivorousDinosaur,
		expectedCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[3],
			PowerStatus: "ACTIVE",
			Type:        "Carnivorous",
		},
		expectedError: nil,
		initialCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[3],
			PowerStatus: "ACTIVE",
			Type:        "Carnivorous",
		},
	}, { // failure to add to cage at max capacity
		description:   testAddDescriptions[4],
		dinosaurToAdd: testFirstCarnivorousDinosaur,
		expectedCage: Cage{
			Capacity: 1,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[4],
			PowerStatus: "ACTIVE",
			Type:        "Carnivorous",
		},
		expectedError: errors.New(errPrefix + errMaxCapacity),
		initialCage: Cage{
			Capacity: 1,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[4],
			PowerStatus: "ACTIVE",
			Type:        "Carnivorous",
		},
	}, { // failure add carnivorous dinosaur to existing carnivorous cage with unlike species
		description:   testAddDescriptions[5],
		dinosaurToAdd: testSecondCarnivorousDinosaur,
		expectedCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[5],
			PowerStatus: "ACTIVE",
			Type:        "Carnivorous",
		},
		expectedError: errors.New(errPrefix + errUnlikeCarnivores),
		initialCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[5],
			PowerStatus: "ACTIVE",
			Type:        "Carnivorous",
		},
	}, { // successful add herbivorous dinosaur to existing herbivorous cage with like species
		description:   testAddDescriptions[6],
		dinosaurToAdd: testFirstHerbivorousDinosaur,
		expectedCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstHerbivorousDinosaur,
				testFirstHerbivorousDinosaur,
			},
			Name:        testAddDescriptions[6],
			PowerStatus: "ACTIVE",
			Type:        "Herbivorous",
		},
		expectedError: nil,
		initialCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstHerbivorousDinosaur,
			},
			Name:        testAddDescriptions[6],
			PowerStatus: "ACTIVE",
			Type:        "Herbivorous",
		},
	}, { //successful add herbivorous dinosaur to existing herbivorous cage with unlike species
		description:   testAddDescriptions[7],
		dinosaurToAdd: testSecondHerbivorousDinosaur,
		expectedCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstHerbivorousDinosaur,
				testSecondHerbivorousDinosaur,
			},
			Name:        testAddDescriptions[7],
			PowerStatus: "ACTIVE",
			Type:        "Herbivorous",
		},
		expectedError: nil,
		initialCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstHerbivorousDinosaur,
			},
			Name:        testAddDescriptions[7],
			PowerStatus: "ACTIVE",
			Type:        "Herbivorous",
		},
	}, { // failure to add carnivorous dinosaur to existing herbivorous cage
		description:   testAddDescriptions[8],
		dinosaurToAdd: testFirstCarnivorousDinosaur,
		expectedCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstHerbivorousDinosaur,
			},
			Name:        testAddDescriptions[8],
			PowerStatus: "ACTIVE",
			Type:        "Herbivorous",
		},
		expectedError: errors.New(errPrefix + errAddCarnivoreToHerbivorous),
		initialCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstHerbivorousDinosaur,
			},
			Name:        testAddDescriptions[8],
			PowerStatus: "ACTIVE",
			Type:        "Herbivorous",
		},
	}, { // failure to add herbivorous dinosaur to existing carnivorous cage
		description:   testAddDescriptions[9],
		dinosaurToAdd: testFirstHerbivorousDinosaur,
		expectedCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[9],
			PowerStatus: "ACTIVE",
			Type:        "Carnivorous",
		},
		expectedError: errors.New(errPrefix + errAddHerbivoreToCarnivorous),
		initialCage: Cage{
			Capacity: 2,
			Dinosaurs: []dinosaur.Dinosaur{
				testFirstCarnivorousDinosaur,
			},
			Name:        testAddDescriptions[9],
			PowerStatus: "ACTIVE",
			Type:        "Carnivorous",
		},
	}}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			var park Park
			// setup existing cage for test
			if test.initialCage.Name != "" {
				park.addCage(test.initialCage.Name, test.initialCage.Capacity)
				if test.initialCage.PowerStatus == "ACTIVE" {
					park.togglePower(test.initialCage.Name)
				}
				if len(test.initialCage.Dinosaurs) > 0 {
					for _, dino := range test.initialCage.Dinosaurs {
						park.addDinosaurToCage(test.initialCage.Name, dino.Name, dino.Species)
					}
				}
			}

			err := park.addDinosaurToCage(test.initialCage.Name, test.dinosaurToAdd.Name, test.dinosaurToAdd.Species)
			assert.Equal(t, test.expectedCage, park.cages[test.initialCage.Name])
			assert.Equal(t, test.expectedError, err)
		})
	}
}

// func TestAddCage(t *testing.T) {
// 	tests := []struct {
// 		description   string
// 		expectedError error
// 		cageName      string
// 		existingCages []Cage
// 	}{{ // successful add cage
// 		description:   testAddDescriptions[0],
// 		expectedCage: Cage{
// 			Capacity: 1,
// 			Dinosaurs: []dinosaur.Dinosaur{
// 			},
// 			Name:        "successful add cage",
// 			PowerStatus: "ACTIVE",
// 			Type:        "carnivorous",
// 		},
// 		expectedError: nil,
// 		existingCages: []Cage{},
// 		},
// 	}}

// 	for _, test := range tests {
// 		t.Run(test.description, func(t *testing.T) {

// 			err := AddDinosaurToCage(test.initialCage.Name, test.dinosaurToAdd)
// 			assert.Equal(t, test.expectedCage, test.initialCage)
// 			if assert.Error(t, err) {
// 				assert.Equal(t, test.expectedError, err)
// 			}
// 		})
// 	}
// }
