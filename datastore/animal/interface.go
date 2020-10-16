package animal

import "webapp/entities"

// IAnimalRepository ...
type IAnimalRepository interface {
	Get(id int) ([]entities.Animal, error)
	Create(entities.Animal) (entities.Animal, error)
}
