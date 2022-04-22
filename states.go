package main

import "fmt"

// chapter four
type State struct {
	name       string
	population int
	capital    string
}

type States map[string]State

func populateMapOfStats() {
	states := make(States, 6)

	states["GO"] = State{"Goiás", 6434052, "Goiânia"}
	states["PB"] = State{"Paraíba", 3914418, "João Pessoa"}
	states["PR"] = State{"Paraná", 10997462, "Curitiba"}
	states["RN"] = State{"Rio Grande do Norte", 3373960, "Natal"}
	states["AM"] = State{"Amazonas", 3807923, "Manaus"}
	states["SE"] = State{"Sergipe", 2228489, "Aracaju"}

	states.Print()
}

func (states States) lookUp(key string) {
	if states.HasKey(key) {
		fmt.Println(states[key])
	}
}

func (states States) HasKey(key string) bool {
	_, hasKey := states[key]
	return hasKey
}

func (states States) Print() {
	fmt.Println(states)
}

func (states States) Delete(key string) {
	delete(states, key)
}

func (states States) IterateStates() {
	for initials, state := range states {
		fmt.Printf("%s (%s) has %d habitants.\n", state.name, initials, state.population)
	}
}
