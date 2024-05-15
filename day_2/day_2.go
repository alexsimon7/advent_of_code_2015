package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//import "slices"

func main() {
	file, err := os.ReadFile("./present_list.txt")

	if err != nil {
		log.Panic(err)
	}

	answer := 0
	answer_ribbon := 0
	current_present := make([]string, 3)
	running_dim := ""
	current_ele := 0

	for _, runeValue := range file {
		if runeValue == 'x' {

			current_present[current_ele] = running_dim
			current_ele += 1
			running_dim = ""
		} else if runeValue == '\n' {

			current_present[current_ele] = running_dim
			fmt.Println(current_present)

			length, err := strconv.Atoi(current_present[0])
			width, err2 := strconv.Atoi(current_present[1])
			height, err3 := strconv.Atoi(current_present[2])
			ribbon_needed := 0

			if err != nil || err2 != nil || err3 != nil {
				panic(err)
			}

			surface_area_1 := (2 * length * width)
			surface_area_2 := (2 * width * height)
			surface_area_3 := (2 * height * length)

			total_surface_area := surface_area_1 + surface_area_2 + surface_area_3

			if surface_area_1 <= surface_area_2 && surface_area_1 <= surface_area_3 {
				total_surface_area += surface_area_1 / 2
			} else if surface_area_2 <= surface_area_1 && surface_area_2 <= surface_area_3 {
				total_surface_area += surface_area_2 / 2
			} else {
				total_surface_area += surface_area_3 / 2
			}

			if length <= width && length <= height {
				ribbon_needed += 2 * length

				if width <= height {
					ribbon_needed += 2 * width
				} else {
					ribbon_needed += 2 * height
				}
			} else if width <= length && width <= height {
				ribbon_needed += 2 * width

				if length <= height {
					ribbon_needed += 2 * length
				} else {
					ribbon_needed += 2 * height
				}
			} else {
				ribbon_needed += 2 * height

				if length <= width {
					ribbon_needed += 2 * length
				} else {
					ribbon_needed += 2 * width
				}
			}

			answer_ribbon += (ribbon_needed + (length * width * height))

			answer += total_surface_area

			current_present = make([]string, 3)
			current_ele = 0
			running_dim = ""
		} else {
			running_dim += string(runeValue)
		}
	}

	fmt.Println(answer_ribbon)
}
