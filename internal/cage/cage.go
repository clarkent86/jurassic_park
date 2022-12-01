package cage

import (
	"errors"

	"github.com/clarkent86/jurassic_park/internal/dinosaur"
)

type Cage struct {
	Capacity    int
	Dinosaurs   []dinosaur.Dinosaur
	Name        string
	PowerStatus string
	Type        string
}

func (cage *Cage) AddDinosaurToCage(dinosaur.Dinosaur) error {
	return errors.New("placeholder for TDD")
}
