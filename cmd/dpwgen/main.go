package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strconv"

	"github.com/tthanh/dpwgen/wordlist"
)

func main() {
	n := flag.Int("n", 6, "number of words")
	flag.Parse()
	words := []string{}
	for w := 0; w < *n; w++ {
		k := ""
		for i := 0; i < 5; i++ {
			m, err := rand.Int(rand.Reader, big.NewInt(int64(6)))
			if err != nil {
				panic(err)
			}
			k += strconv.Itoa(int(m.Int64()) + 1)
		}
		words = append(words, k)
	}

	for _, v := range words {
		fmt.Printf("%s - %s\n", v, wordlist.EFFLarge[v])
	}
}
