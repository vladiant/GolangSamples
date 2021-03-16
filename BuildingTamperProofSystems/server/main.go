package main

import (
	"log"

	"github.com/vladiant/GolangSamples/BuildingTamperProofSystems/internal/immudb"
)

func main() {
	db, err := immudb.NewImmuDB()
	if err != nil {
		log.Fatal("Could not connect to immudb")
	}

	err = db.UpdateSalary("12345", 100000)
	if err != nil {
		log.Fatal("could not update engineers salary")
	}
}
