// repository.go

package repository

import (
	"fmt"
	"sync"

	. "github.com/johansortiz/SuperHero/models"
)

type SuperheroRepository struct {
	superheroes map[string]Superhero
	mutex       sync.RWMutex
}

func NewSuperheroRepository() *SuperheroRepository {
	return &SuperheroRepository{
		superheroes: make(map[string]Superhero),
	}
}

func (r *SuperheroRepository) Add(superhero Superhero) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.superheroes[superhero.Name] = superhero
}

func (r *SuperheroRepository) Get(name string) (Superhero, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	superhero, ok := r.superheroes[name]
	if !ok {
		return Superhero{}, fmt.Errorf("superhero %s not found", name)
	}
	return superhero, nil
}
