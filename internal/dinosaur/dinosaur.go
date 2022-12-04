package dinosaur

import (
	"errors"
	"fmt"
	"strings"
)

var dinoPool = [8]string{"tyrannosaurus", "velociraptor", "spinosaurus", "megalosaurus", "brachiosaurus", "stegosaurus", "ankylosaurus", "triceratops"}

/*
Dinosaur
struct to contain a dinosaur's properties
*/
type Dinosaur struct {
	Name    string
	Species string
}

/*
InitDino
initialized a dino with basic verification
*/
func InitDinosaur(name, species string) (Dinosaur, error) {
	if !isDinoEggAvailable(species) {
		return Dinosaur{}, errors.New("this species of dino does not exist/has has not been synthesized")
	}
	return Dinosaur{Name: name, Species: species}, nil
}

/*
GetDiet
returns the diet of the given dinosaur's species
*/
func (dinosaur Dinosaur) GetDiet() (string, error) {
	switch strings.ToLower(dinosaur.Species) {
	case "tyrannosaurus", "velociraptor", "spinosaurus", "megalosaurus":
		return "Carnivorous", nil
	case "brachiosaurus", "stegosaurus", "ankylosaurus", "triceratops":
		return "Herbivorous", nil
	}
	return "", errors.New("something went wrong determining the dinosaur's diet")
}

/*
isDinoEggAvailable
helper function to quickly return the availability of the given dinosaur species

TODO: compile a full list of dinosaurs to differentiate between existence and
gibberish or misspelled species for a more specific error message
*/
func isDinoEggAvailable(speciesName string) bool {
	for _, v := range dinoPool {
		if v == strings.ToLower(speciesName) {
			return true
		}
	}
	return false
}

func (dinosaur Dinosaur) String() string {
	return fmt.Sprintf("Name: %s\nSpecies: %s", dinosaur.Name, dinosaur.Species)
}
