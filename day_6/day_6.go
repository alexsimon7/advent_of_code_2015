package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 0 is off, 1 is on

	lights_on := 0

	light_array := make([][]int, 1000)

	for i := range light_array {
		light_array[i] = make([]int, 1000)
	}

	// Grab Instructions from Instructions.txt
	instructions, err := os.Open("./directions.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer instructions.Close()

	scanner := bufio.NewScanner(instructions)

	for scanner.Scan() {
		current_line := scanner.Text()
		current_line_slice := strings.Split(current_line, " ")

		current_instruction := make([]int, 4)
		current_method := ""

		//on, off, toggle

		if current_line_slice[0] == "turn" {
			current_method = current_line_slice[1]

			start := current_line_slice[2]
			current_instruction[0], err = strconv.Atoi(start[:strings.Index(start, ",")])
			current_instruction[1], err = strconv.Atoi(start[strings.Index(start, ",")+1:])

			end := current_line_slice[4]

			current_instruction[2], err = strconv.Atoi(end[:strings.Index(end, ",")])
			current_instruction[3], err = strconv.Atoi(end[strings.Index(end, ",")+1:])

		} else {
			current_method = current_line_slice[0]

			start := current_line_slice[1]
			current_instruction[0], err = strconv.Atoi(start[:strings.Index(start, ",")])
			current_instruction[1], err = strconv.Atoi(start[strings.Index(start, ",")+1:])

			end := current_line_slice[3]

			current_instruction[2], err = strconv.Atoi(end[:strings.Index(end, ",")])
			current_instruction[3], err = strconv.Atoi(end[strings.Index(end, ",")+1:])
		}

		for i := current_instruction[0]; i <= current_instruction[2]; i++ {
			for j := current_instruction[1]; j <= current_instruction[3]; j++ {
				if current_method == "on" {
					light_array[i][j] += 1
					lights_on += 1
					//					if light_array[i][j] != 1 {
					//						light_array[i][j] = 1
					//						lights_on += 1
					//					}
				} else if current_method == "off" {
					if light_array[i][j]-1 < 0 {
						light_array[i][j] = 0
					} else {
						light_array[i][j] -= 1
						lights_on -= 1
					}
					//					if light_array[i][j] != 0 {
					//						light_array[i][j] = 0
					//						lights_on -= 1
					//					}
				} else {
					light_array[i][j] += 2
					lights_on += 2
					//					if light_array[i][j] == 1 {
					//						light_array[i][j] = 0
					//						lights_on -= 1
					//					} else {
					//						light_array[i][j] = 1
					//						lights_on += 1
					//					}
				}
			}
		}
	}

	fmt.Printf("Number of lights on: %d \n", lights_on)
}
