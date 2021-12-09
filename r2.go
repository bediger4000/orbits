package main

import (
	"encoding/json"
	"fmt"
	"log"
	"orbits/body"
	"os"
)

func main() {
	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var configuration body.Config

	err = json.Unmarshal(buf, &configuration)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d time steps of %f sec\n", configuration.Steps, configuration.Interval)
	fmt.Printf("%d masses\n", len(configuration.Masses))

	for _, mass := range configuration.Masses {
		fmt.Printf("%s: %f\n", mass.Name, mass.Mass)
	}
}
