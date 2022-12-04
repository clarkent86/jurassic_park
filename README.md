# jurassic_park
a system to keep track of the different cages around the park and the different dinosaurs in each one

## Jurassic Park - Go

### The Problem
It's 1993 and you're the lead software developer for the new Jurassic Park! Park
operations needs a system to keep track of the different cages around the park and the
different dinosaurs in each one. You'll need to develop a JSON formatted RESTful API
to allow the builders to create new cages. It will also allow doctors and scientists the
ability to edit/retrieve the statuses of dinosaurs and cages.

### Business Requirements
Please attempt to implement the following business requirements:
- All requests should respond with the correct HTTP status codes and a response, if
necessary, representing either the success or error conditions.
- Each dinosaur must have a name.
- Each dinosaur is considered an herbivore or a carnivore, depending on its species.
- Carnivores can only be in a cage with other dinosaurs of the same species.
- Each dinosaur must have a species (See enumerated list below, feel free to add
others).
- Herbivores cannot be in the same cage as carnivores.
- Use Carnivore dinosaurs like Tyrannosaurus, Velociraptor, Spinosaurus and
Megalosaurus.
- Use Herbivores like Brachiosaurus, Stegosaurus, Ankylosaurus and Triceratops.

### Technical Requirements
The following technical requirements must be met:
- You are allowed to use scaffolding technology.
- This project should be done with version Golang 1.19 or newer.
- This project can use the Gin API framework or similar technology.
- This should be done using version control, preferably git.
- The project should include a README that addresses anything you may not have
completed. It should also address what additional changes you might need to make
if the application were intended to run in a concurrent environment. Any other
comments or thoughts about the project are also welcome.

### Bonus Points
- Cages have a maximum capacity for how many dinosaurs it can hold.
- Cages know how many dinosaurs are contained.
- Cages have a power status of ACTIVE or DOWN.
- Cages cannot be powered off if they contain dinosaurs.
- Dinosaurs cannot be moved into a cage that is powered down.
- Must be able to query a listing of dinosaurs in a specific cage.
- When querying dinosaurs or cages they should be filterable on their attributes
(Cages on their power status and dinosaurs on species).
- Automated tests that ensure the business logic implemented is correct.

### Other Ideas (outside given prompt)
- Can't have multiple dinosaurs with the same name
- More Diet Specifics (quantity per cage)
- people friendly cages (do they need power? is power an electric fence?)
- Breeding
- Jail break! reassign dinosaurs, tally losses (dino and human??)
- Suggest an available cage when failure to add dinosaur to a cage
- Bulk Add/Delete dinosaurs or create cage with bulk dinosaurs
- Chef Goldblum

### General Thoughts to improve/TODOs
- dockerize
- environment variables to control port & potentially other environmental control
- better format park response
- better http status response codes for errors
- flesh out http response testing (asserting expected error responses & bodies)
- long term park storage (park disappears when app dies)
- authenticate users
- authorize users for cages
- pass around context to allow for better use of authentication/authorization
- reduce path redundancy...don't need add/delete in path if I have proper CRUD methods (which I do now!)
- Swagger docs...using the built in Golang swagger tools to auto generate instead of typing out the Usage section here

### Notes
- Gorilla/Mux was used as it's similar to Gin, but preferred for native support of Golang's internal context package

### Usage

The Jurassic Park app is provided with the bare Golang build/run usage.

```bash
go run cmd/app/app.go
```

The app is started without any cages. To get started, you need to first add a cage, power it on, and you are ready to start adding Dinosaurs!

Add a cage:
```
localhost:8080/add/cage
```
Required query parameters:
```
Capacity: int
cageName: string
```

Remove a cage:
```
localhost:8080/delete/cage
```
Required query parameters:
```
cageName: string
```

Toggle cage power (when cage is empty):
```
localhost:8080/togglePower
```
Required query parameters:
```
cageName: string
```

Add dinosaur:
```
localhost:8080/add/dinosaur
```
Required query parameters:
```
cageName: string
dinosaurName: string
dinosaurSpecies: string
```

Remove a dinosaur from cage:
```
localhost:8080/delete/dinosaur
```
Required query parameters:
```
cageName: string
dinosaurName: string
dinosaurSpecies: string
```

Check park status:
```
localhost:8080/park/status
```
No required query parameters