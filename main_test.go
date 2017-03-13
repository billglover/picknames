package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func BenchmarkPick1(b *testing.B)  { pickRandomN(1, b) }
func BenchmarkPick5(b *testing.B)  { pickRandomN(1, b) }
func BenchmarkPick10(b *testing.B) { pickRandomN(1, b) }

func pickRandomN(c int, b *testing.B) {
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
