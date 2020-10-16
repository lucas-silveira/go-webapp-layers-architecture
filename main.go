package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	AnimalRepository "webapp/datastore/animal"
	HttpAnimal "webapp/delivery/animal"
	"webapp/driver"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	config := driver.MySQLConfig{
		Host:     os.Getenv("SQL_HOST"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port:     os.Getenv("SQL_PORT"),
		Db:       os.Getenv("SQL_DB"),
	}

	var err error

	db, err := driver.ConnectToMySQL(config)

	if err != nil {
		log.Println("Could not connect to sql, err:", err)
		return
	}

	animalRepository := AnimalRepository.New(db)
	httpAnimal := HttpAnimal.New(animalRepository)

	http.HandleFunc("/animal", httpAnimal.Handler)
	fmt.Println(http.ListenAndServe(":3000", nil))
}
