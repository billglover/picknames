package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	count = flag.Int("count", 10, "the number of names to generate")
)

func main() {
	flag.Parse()

	f, err := os.Create("test_data.txt")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	for n := 1; n <= *count; n++ {
		fmt.Fprintf(f, "First Last %06d\n", n)
	}
}
