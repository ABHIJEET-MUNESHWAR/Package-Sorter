package main

import (
	"fmt"
	"github.com/ABHIJEET-MUNESHWAR/Package-Sorter/sorter"
)

func main() {
	packages := []struct {
		width, height, length, mass int
	}{
		{140, 100, 10, 10},  // STANDARD
		{160, 120, 90, 19},  // SPECIAL
		{200, 160, 160, 30}, // REJECTED
		{100, 100, 10, 5},   // STANDARD
	}
	results := make(chan string, len(packages))

	for _, pkg := range packages {
		go func(w, h, l, m int) {
			results <- sorter.Sort(w, h, l, m)
		}(pkg.width, pkg.height, pkg.length, pkg.mass)
	}
	for range packages {
		fmt.Println(<-results)
	}
}