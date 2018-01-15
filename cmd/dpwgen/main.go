package main

import (
	"fmt"

	"github.com/tthanh/dpwgen/wordlist"
)

func main() {
	for _, s := range wordlist.EFFLarge {
		fmt.Println(s)
	}
}
