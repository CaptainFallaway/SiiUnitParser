package main

import (
	"fmt"
	"os"

	"github.com/CaptainFallaway/SiiUnitParser/pkg/siiunit"
)

func main() {
	file, err := os.Open("data/game-decoded.sii")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	units, err := siiunit.ParseAllUnits(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Parsed", len(units), "units")

	for _, unit := range units {
		fmt.Println(unit)
	}
}
