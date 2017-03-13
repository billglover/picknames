package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	count = flag.Int("n", 1, "the number of names to pick from the list")
)

type person struct {
	Name  string
	Score int
}

type byScore []person

func (a byScore) Len() int           { return len(a) }
func (a byScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	// open the list of names
	// TODO: should be able to accept STDIN as well
	f, err := os.Open("data/test_data.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	// read the names into memory and assign each a random score
	r := bufio.NewReader(f)

	// pick names from the list
	results, _ := pickRandom(*count, r)

	// return the number of names requested
	for c := 0; c < *count; c++ {
		log.Printf("%s\n", results[c])
	}

}

func pickRandom(n int, r *bufio.Reader) (p []string, e error) {
	p = make([]string, n)

	var list []person

	for {
		name, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		p := person{
			Name:  strings.TrimRight(name, "\n"),
			Score: rand.Int(),
		}
		list = append(list, p)

	}

	// sort the list in ascending order
	sort.Sort(byScore(list))

	// return the number of names requested
	for c := 0; c < *count; c++ {
		p[c] = list[c].Name
	}

	return
}
