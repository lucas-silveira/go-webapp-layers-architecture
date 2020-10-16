package datastore

import "webapp/entities"

// Animal ...
type Animal interface {
	Get(id int) ([]entities.Animal, error)
	Create(entities.Animal) (entities.Animal, error)
}
