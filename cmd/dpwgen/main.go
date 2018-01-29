package main

import (
	"crypto/rand"
	"flag"
	"math/big"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/tthanh/dpwgen/wordlist"
)

func main() {
	n := flag.Int("n", 6, "number of words")
	flag.Parse()
	keys := []string{}
	for w := 0; w < *n; w++ {
		k := ""
		for i := 0; i < 5; i++ {
			m, err := rand.Int(rand.Reader, big.NewInt(int64(6)))
			if err != nil {
				panic(err)
			}
			k += strconv.Itoa(int(m.Int64()) + 1)
		}
		keys = append(keys, k)
	}

	data := [][]string{
		[]string{},
	}

	for _, v := range keys {
		data[0] = append(data[0], wordlist.EFFLarge[v])
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(keys)

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}
