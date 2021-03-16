package main

import (
	"log"

	"github.com/vladiant/GolangSamples/BuildingTamperProofSystems/internal/immudb"
)

func main() {
	db, err := immudb.NewImmuDB()
	if err != nil {
		log.Fatal(err)
	}

	salary, err := db.GetVerifiedSalary("12345")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Salary Information: %d\n", salary)
}
