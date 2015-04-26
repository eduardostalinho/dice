package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/justinian/dice"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	rollDesc := strings.Join(os.Args[1:], " ")
	res, err := dice.Roll(rollDesc)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	} else {
		fmt.Println(res)
	}
}