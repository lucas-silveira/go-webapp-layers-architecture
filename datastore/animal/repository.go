package animal

import (
	"database/sql"
	"webapp/entities"
)

// Repository ...
type Repository struct {
	db *sql.DB
}

// New ...
func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Get ..;
func (animalRepo Repository) Get(id int) ([]entities.Animal, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if id != 0 {
		rows, err = animalRepo.db.Query("SELECT * FROM animals WHERE id = ?", id)
	} else {
		rows, err = animalRepo.db.Query("SELECT * FROM animals")
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var animals []entities.Animal

	for rows.Next() {
		var a entities.Animal
		_ = rows.Scan(&a.ID, &a.Name, &a.Age)
		animals = append(animals, a)
	}

	return animals, nil
}

// Create ...
func (animalRepo Repository) Create(animal entities.Animal) (entities.Animal, error) {
	res, err := animalRepo.db.Exec("INSERT INTO animals (name, age) VALUES(?, ?)", animal.Name, animal.Age)

	if err != nil {
		return entities.Animal{}, err
	}

	id, _ := res.LastInsertId()
	animal.ID = int(id)

	return animal, nil
}
