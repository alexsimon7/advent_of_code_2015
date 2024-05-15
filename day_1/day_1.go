package main

import "fmt"
import "os"
import "log"

func main() {
	instructions, err := os.ReadFile("./instructions.txt")

	if err != nil {
		log.Fatal(err)
	}

	answer := 0

	for idx, runeValue := range instructions {
		if runeValue == '(' {
			answer += 1
		} else if runeValue == ')' {
			answer -= 1
		}

		if answer == -1 {
			fmt.Println(idx + 1)
			break
		}
	}

}
