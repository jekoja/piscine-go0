package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 || args[0] != "-c" {
		os.Exit(1)
	}

	var n int
	if _, err := fmt.Sscanf(args[1], "%d", &n); err != nil || n < 0 {
		os.Exit(1)
	}

	files := args[2:]
	exitCode := 0
	multiple := len(files) > 1

	for i, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			// Print the error exactly as Go provides it (no prefix)
			fmt.Println(err)
			exitCode = 1
			continue
		}

		if multiple {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", file)
		}

		start := len(data) - n
		if start < 0 {
			start = 0
		}
		fmt.Printf("%s", string(data[start:]))
	}

	if exitCode != 0 {
		os.Exit(1)
	}
}
