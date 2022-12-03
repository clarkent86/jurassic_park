package dinosaur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDinosaur(t *testing.T) {
	tests := []struct {
		description      string
		dinosaurToAdd    Dinosaur
		expectedCreation Dinosaur
	}{{
		description:      "initialize Tyrannosaurus",
		dinosaurToAdd:    Dinosaur{"T-Rex", "Tyrannosaurus"},
		expectedCreation: Dinosaur{Name: "T-Rex", Species: "Tyrannosaurus"},
	}, {
		description:      "initialize Velociraptor lower-case",
		dinosaurToAdd:    Dinosaur{"V-Rap", "Velociraptor"},
		expectedCreation: Dinosaur{Name: "V-Rap", Species: "Velociraptor"},
	}, {
		description:      "initialize Spinosaurus upper-case",
		dinosaurToAdd:    Dinosaur{"S-aurus", "SPINOSAURUS"},
		expectedCreation: Dinosaur{Name: "S-aurus", Species: "SPINOSAURUS"},
	}, {
		description:      "initialize Megalosaurus mixed-case",
		dinosaurToAdd:    Dinosaur{"M-aurus", "MeGaLoSaUrUs"},
		expectedCreation: Dinosaur{Name: "M-aurus", Species: "MeGaLoSaUrUs"},
	}, {
		description:      "initialize Brachiosaurus",
		dinosaurToAdd:    Dinosaur{"B-Saurus", "Brachiosaurus"},
		expectedCreation: Dinosaur{Name: "B-Saurus", Species: "Brachiosaurus"},
	}, {
		description:      "initialize Stegosaurus same name as species",
		dinosaurToAdd:    Dinosaur{"Stegosaurus", "Stegosaurus"},
		expectedCreation: Dinosaur{Name: "Stegosaurus", Species: "Stegosaurus"},
	}, {
		description:      "initialize Ankylosaurus",
		dinosaurToAdd:    Dinosaur{"A-saurus", "Ankylosaurus"},
		expectedCreation: Dinosaur{Name: "A-saurus", Species: "Ankylosaurus"},
	}, {
		description:      "initialize Triceratops",
		dinosaurToAdd:    Dinosaur{"T-tops", "Triceratops"},
		expectedCreation: Dinosaur{Name: "T-tops", Species: "Triceratops"},
	}, {
		description:      "mispelling",
		dinosaurToAdd:    Dinosaur{"T-tops", "Tricerratops"},
		expectedCreation: Dinosaur{},
	}, {
		description:      "non-synthesised dinosaur",
		dinosaurToAdd:    Dinosaur{"Crandios", "Fukuiraptor"},
		expectedCreation: Dinosaur{},
	}}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			dinosaur, _ := InitDinosaur(test.dinosaurToAdd.Name, test.dinosaurToAdd.Species)
			assert.Equal(t, test.expectedCreation, dinosaur)
		})
	}
}
