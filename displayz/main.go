package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main () {
	args := os.Args [1:] 

	if len(args) < 2 {
		if len(args) == 1 {
			word := args[0]
			for _, r := range word {
				if r == 'z' {
					z01.PrintRune('z')
					z01.PrintRune('\n')
					return
				}
			}
		}
		z01.PrintRune('z')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('z')
		z01.PrintRune('\n')
	}
}

		