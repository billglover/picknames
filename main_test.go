package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func Benchmark1Pick1(b *testing.B)  { pickRandomN1(1, b) }
func Benchmark1Pick5(b *testing.B)  { pickRandomN1(1, b) }
func Benchmark1Pick10(b *testing.B) { pickRandomN1(1, b) }

func Benchmark2Pick1(b *testing.B)  { pickRandomN2(1, b) }
func Benchmark2Pick5(b *testing.B)  { pickRandomN2(1, b) }
func Benchmark2Pick10(b *testing.B) { pickRandomN2(1, b) }

func pickRandomN1(c int, b *testing.B) {
	f, err := os.Open("data/test_data.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = pickRandom(c, r)
		f.Seek(0, 0)
	}
}

func pickRandomN2(c int, b *testing.B) {
	f, err := os.Open("data/test_data.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = pickRandom2(c, f)
		f.Seek(0, 0)
	}
}
