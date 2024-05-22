package main

import "os"
import "fmt"

import "strings"

//import "maps"

const vowels string = "aeiou"

func main() {
	list, err := os.ReadFile("./santa_list.txt")

	if err != nil {
		panic(err)
	}

	var previous_previous_previous_char string
	var previous_previous_char string
	var previous_char string
	nice_strings := 0

	pairs := make(map[string]int)

	//	three_vowels := 0
	//	double_letter := false
	//	no_spec_string := true

	one_letter_between := false
	repeating_pair := false

	for _, runeValue := range list {

		current_char := string(runeValue)
		//		if strings.Contains(vowels, current_char) {
		//			three_vowels += 1
		//		}

		//		if previous_char == current_char {
		//			double_letter = true
		//		}

		//		if (previous_char == "a" && current_char == "b") || (previous_char == "c" && current_char == "d") || (previous_char == "p" && current_char == "q") || (previous_char == "x" && current_char == "y") {
		//			no_spec_string = false
		//		}

		if previous_previous_char == current_char {
			one_letter_between = true
			fmt.Printf("Found One Letter Between %s %s %s \n", previous_previous_char, previous_char, current_char)
		}

		if previous_char != "" && current_char != "\n" {
			current_pair := previous_char + current_char

			if previous_previous_char != previous_char || previous_char != current_char {
				if pairs[current_pair] != 0 {
					pairs[current_pair] += 1

					if pairs[current_pair] == 2 {
						repeating_pair = true
					}

				} else {
					pairs[current_pair] = 1
				}
			} else if previous_previous_previous_char == previous_previous_char && previous_previous_char == previous_char && previous_char == current_char {
				if pairs[current_pair] != 0 {
					pairs[current_pair] += 1

					if pairs[current_pair] == 2 {
						repeating_pair = true
					}
				} else {
					pairs[current_pair] = 1
				}
			}
		}

		if current_char == "\n" {
			//			if three_vowels >= 3 && double_letter && no_spec_string {
			//				nice_strings += 1
			//			}

			if one_letter_between && repeating_pair {
				nice_strings += 1
				fmt.Println("Nice")
			}

			//			three_vowels = 0
			//			double_letter = false
			//			no_spec_string = true

			fmt.Println(pairs)

			one_letter_between = false
			repeating_pair = false
			previous_previous_previous_char = ""
			previous_previous_char = ""
			previous_char = ""
			current_char = ""
			pairs = make(map[string]int)
			continue
		}
		previous_previous_previous_char = strings.Clone(previous_previous_char)
		previous_previous_char = strings.Clone(previous_char)
		previous_char = strings.Clone(current_char)
	}

	fmt.Println(nice_strings)

}
