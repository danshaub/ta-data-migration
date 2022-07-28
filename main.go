package main

import (
	"log"

	"github.com/danshaub/ta-data-migration/csvfile"
)

func main() {
	path := "applicants.csv"
	log.Printf("Finding csv at %s", path)

	apps := ReadFile(path)

	log.Print(apps)
}
