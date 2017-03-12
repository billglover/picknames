package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data/test_data.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)
	for {
		name, err := r.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Print(name)
	}
}
