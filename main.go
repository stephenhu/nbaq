package main

import (
	"flag"
	"fmt"
	//"log"
)


var (
  src 			string
	cache			NbaCache
)


func initFlags() {
  flag.StringVar(&src, "src", CURRENT_DIR, "source directory")
} // initFlags


func main() {

	initFlags()

	flag.Parse()

	fmt.Printf("Starting %s...\n", version())

	cache = NbaCache{
		Seasons: map[string]Season{},
	}

	initCache()

} // main
