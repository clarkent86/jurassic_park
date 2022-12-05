package cage

import (
	"errors"

	"github.com/clarkent86/jurassic_park/internal/dinosaur"
)

const errPrefix = "[internal][cage] "
const errDownCage = "cannot add a dinosaur to a down cage"
const errMaxCapacity = "cannot add dinosaur to a cage at max capacity"
const errUnlikeCarnivores = "cannot add unlike carnivorous dinosaur species together within the same cage"
const errAddCarnivoreToHerbivorous = "cannot add a carnivorous dinosaur to an herbivorous cage"
const errAddHerbivoreToCarnivorous = "cannot add an herbivorous dinosaur to a carnivorous cage"
const errNonExistentCage = "cage does not exist"
const errCageAlreadyExists = "cage already exists"
const errPowerCageWithDinosaurs = "cannot power down a cage with dinosaurs"
const errRemoveNonEmptyCage = "cannot remove cage with dinosaurs"
const errRemovePoweredCage = "cannot remove powered cage"
const errDinosaurDoesNotExist = "dinosaur does not exist"

type Park struct {
	cages map[string]Cage
}

type Cage struct {
	Capacity    int
	Dinosaurs   []dinosaur.Dinosaur
	Name        string
	PowerStatus string
	Type        string
}

func (park *Park) addDinosaurToCage(cageName string, dinosaurName, dinosaurSpecies string) error {
	if _, found := park.cages[cageName]; !found {
		return errors.New(errPrefix + errNonExistentCage)
	}
	cage := park.cages[cageName]
	if cage.Capacity == len(cage.Dinosaurs) {
		return errors.New(errPrefix + errMaxCapacity)
	}

	dinosaur, err := dinosaur.InitDinosaur(dinosaurName, dinosaurSpecies)
	if err != nil {
		return err
	}
	dinosaurDiet, err := dinosaur.GetDiet()
	if err != nil {
		return err
	}

	if len(cage.Dinosaurs) == 0 {
		if cage.PowerStatus == "DOWN" {
			return errors.New(errPrefix + errDownCage)
		}
		cage.Dinosaurs = append(cage.Dinosaurs, dinosaur)
		cage.Type = dinosaurDiet
		park.cages[cageName] = cage
		return nil
	}

	if cage.Type != dinosaurDiet {
		if cage.Type == "carnivorous" {
			return errors.New(errPrefix + errAddHerbivoreToCarnivorous)
		}
		return errors.New(errPrefix + errAddCarnivoreToHerbivorous)
	}

	// diets will be equal since we passed the last conditional
	if cage.Type == "herbivorous" {
		cage.Dinosaurs = append(cage.Dinosaurs, dinosaur)
		park.cages[cageName] = cage
		return nil
	}

	// only diet possible is carnivorous, cage will have all the same species of first dino in cage
	if cage.Dinosaurs[0].Species != dinosaur.Species {
		return errors.New(errPrefix + errUnlikeCarnivores)
	}

	cage.Dinosaurs = append(cage.Dinosaurs, dinosaur)
	park.cages[cageName] = cage

	return nil
}

func (park *Park) addCage(cageName string, capacity int) error {
	if _, found := park.cages[cageName]; found {
		return errors.New(errPrefix + errCageAlreadyExists)
	}
	// TODO: Make this a sync.Once()
	if park.cages == nil {
		park.cages = make(map[string]Cage)
	}
	park.cages[cageName] = Cage{
		Capacity:    capacity,
		Dinosaurs:   []dinosaur.Dinosaur{},
		Name:        cageName,
		PowerStatus: "DOWN",
		Type:        "",
	}
	return nil
}

func (park *Park) removeCage(cageName string) error {
	if _, found := park.cages[cageName]; !found {
		return errors.New(errPrefix+errNonExistentCage)
	}
	cage := park.cages[cageName]; {
		if len(cage.Dinosaurs) > 0 {
			return errors.New(errPrefix+ errRemoveNonEmptyCage)
		}
		if cage.PowerStatus == "ACTIVE"{
			return errors.New(errPrefix+ errRemovePoweredCage)
		}
	}
	delete(park.cages[cageName])
	return nil
}

func (park *Park) removeDinosaurFromCage(cageName, dinosaurName, dinosaurSpecies string) error {
	if _, found := park.cages[cageName]; !found {
		return errors.New(errPrefix+errNonExistentCage)
	}
	cage := park.cages[cageName]
	for i, dinosaur := range cage.Dinosaurs{
		if dinosaur.Name == dinosaurName && dinosaur.Species == dinosaurSpecies {
			delete(park.cages[cageName].Dinosaurs[i])
		}
	}
	return errors.New(errPrefix+errDinosaurDoesNotExist)
}

func (park *Park) togglePower(cageName string) error {
	if _, found := park.cages[cageName]; !found {
		return errors.New(errPrefix + errNonExistentCage)
	}
	cage := park.cages[cageName]
	if cage.PowerStatus == "ACTIVE" {
		if len(cage.Dinosaurs) > 0 {
			return errors.New(errPrefix + errPowerCageWithDinosaurs)
		}
		park.cages[cageName] = Cage{
			Capacity:    cage.Capacity,
			Dinosaurs:   cage.Dinosaurs,
			Name:        cage.Name,
			PowerStatus: "DOWN",
			Type:        cage.Type,
		}
		return nil
	}
	park.cages[cageName] = Cage{
		Capacity:    cage.Capacity,
		Dinosaurs:   cage.Dinosaurs,
		Name:        cage.Name,
		PowerStatus: "ACTIVE",
		Type:        cage.Type,
	}
	return nil
}
