package main

import (
	"fmt"
	"os"
)

func main() {
	gt, err := readGoogleTrends("CH")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, s := range gt {
		fmt.Printf("Trending %s\n", s)
	}
}
