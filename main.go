package main

import (
	"log"
)

func main() {
	path := "applicants.csv"
	log.Printf("Finding csv at %s", path)

	apps := ReadFile(path)

	log.Print(apps)
}
