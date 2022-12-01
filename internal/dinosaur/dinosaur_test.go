package dinosaur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDinosaur(t *testing.T) {
	tests := []struct {
		description      string
		name             string
		species          string
		expectedCreation Dinosaur
	}{{
		description:      "initialize Tyrannosaurus",
		name:             "T-Rex",
		species:          "Tyrannosaurus",
		expectedCreation: Dinosaur{Name: "T-Rex", Species: "Tyrannosaurus"},
	}, {
		description:      "initialize Velociraptor lower-case",
		name:             "V-Rap",
		species:          "Velociraptor",
		expectedCreation: Dinosaur{Name: "V-Rap", Species: "Velociraptor"},
	}, {
		description:      "initialize Spinosaurus upper-case",
		name:             "S-aurus",
		species:          "SPINOSAURUS",
		expectedCreation: Dinosaur{Name: "S-aurus", Species: "SPINOSAURUS"},
	}, {
		description:      "initialize Megalosaurus mixed-case",
		name:             "M-aurus",
		species:          "MeGaLoSaUrUs",
		expectedCreation: Dinosaur{Name: "M-aurus", Species: "MeGaLoSaUrUs"},
	}, {
		description:      "initialize Brachiosaurus",
		name:             "B-Saurus",
		species:          "Brachiosaurus",
		expectedCreation: Dinosaur{Name: "B-Saurus", Species: "Brachiosaurus"},
	}, {
		description:      "initialize Stegosaurus same name as species",
		name:             "Stegosaurus",
		species:          "Stegosaurus",
		expectedCreation: Dinosaur{Name: "Stegosaurus", Species: "Stegosaurus"},
	}, {
		description:      "initialize Ankylosaurus",
		name:             "A-saurus",
		species:          "Ankylosaurus",
		expectedCreation: Dinosaur{Name: "A-saurus", Species: "Ankylosaurus"},
	}, {
		description:      "initialize Triceratops",
		name:             "T-tops",
		species:          "Triceratops",
		expectedCreation: Dinosaur{Name: "T-tops", Species: "Triceratops"},
	}, {
		description:      "mispelling",
		name:             "T-tops",
		species:          "Tricerratops",
		expectedCreation: Dinosaur{},
	}, {
		description:      "non-synthesised dinosaur",
		name:             "Crandios",
		species:          "Fukuiraptor",
		expectedCreation: Dinosaur{},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dinosaur, _ := InitDinosaur(test.name, test.species)
			assert.Equal(t, test.expectedCreation, dinosaur)
		})
	}
}
