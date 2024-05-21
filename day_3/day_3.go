package main

import "fmt"

import "os"
import "log"
import "slices"

func main() {
	route, err := os.ReadFile("./delivery_route.txt")

	if err != nil {
		log.Fatal(err)
	}

	santa_x_axis := 0
	santa_y_axis := 0
	robo_x_axis := 0
	robo_y_axis := 0

	houses_visited := 0

	paths_taken := make(map[int][]int)

	paths_taken[santa_x_axis] = append(paths_taken[santa_x_axis], santa_y_axis)
	houses_visited += 1

	for idx, runeValue := range route {

		if idx%2 == 0 {
			if runeValue == '>' {
				santa_x_axis += 1
			} else if runeValue == '<' {
				santa_x_axis -= 1
			} else if runeValue == '^' {
				santa_y_axis += 1
			} else if runeValue == 'v' {
				santa_y_axis -= 1
			}

			if !slices.Contains(paths_taken[santa_x_axis], santa_y_axis) {
				houses_visited += 1
				paths_taken[santa_x_axis] = append(paths_taken[santa_x_axis], santa_y_axis)
			}

		} else {
			if runeValue == '>' {
				robo_x_axis += 1
			} else if runeValue == '<' {
				robo_x_axis -= 1
			} else if runeValue == '^' {
				robo_y_axis += 1
			} else if runeValue == 'v' {
				robo_y_axis -= 1
			}

			if !slices.Contains(paths_taken[robo_x_axis], robo_y_axis) {
				houses_visited += 1
				paths_taken[robo_x_axis] = append(paths_taken[robo_x_axis], robo_y_axis)
			}
		}

	}

	fmt.Println("Houses visited: ", houses_visited)
}
