package main

import (
	"fmt"
	"os"
	"time"

	"github.com/CaptainFallaway/SiiUnitParser/pkg/siiunit"
)

func main() {
	file, err := os.Open("data/game-decoded.sii")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	start := time.Now()

	units, err := siiunit.ParseAllUnitsConcurrent(file, siiunit.OptWorkerCount(8))
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start)

	for _, unit := range units {
		fmt.Println(unit)
	}

	unit := units[0]

	unlockedDealersAttr, _ := unit.Attrs.Get("unlocked_dealers")

	unlockedDealersAttrs, err := unlockedDealersAttr.Arr()
	if err != nil {
		panic(err)
	}

	var unlockedDealers []string
	for _, attr := range unlockedDealersAttrs {
		strVal, err := attr.String()
		if err != nil {
			panic(err)
		}
		unlockedDealers = append(unlockedDealers, strVal)
	}

	fmt.Println(unlockedDealers)

	fmt.Println("Parsed", len(units), "units")
	fmt.Println("Parsing took", elapsed)
}
