package main

import (
	"fmt"
	"get-aoe2-civ/internal/getaoe2civ"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please supply a config file")
		os.Exit(1)
	}

	getaoe2civ.GetCiv(os.Args[1])
}
