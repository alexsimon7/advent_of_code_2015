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
	circuit := make(map[string]uint16)

	var instruction []string
	operator := ""
	operand_one := ""
	operand_two := ""
	destination := ""

	file, err := os.Open("./test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		current_direction := scanner.Text()

		current_direction_slice := strings.Split(current_direction, "->")

		instruction = strings.Split(current_direction_slice[0], " ")
		destination = strings.TrimSpace(current_direction_slice[1])

		if len(instruction) == 2 {
			operand_one = instruction[0]

			if _, err := strconv.Atoi(operand_one); err == nil {
				value, _ := strconv.ParseInt(operand_one, 10, 16)
				circuit[destination] = uint16(value)
			} else {
				circuit[destination] = circuit[operand_one]
			}

		} else if len(instruction) == 3 {
			operator = instruction[0]
			operand_one = instruction[1]
			circuit[destination] = ^uint16(circuit[operand_one])

		} else {
			operand_one = instruction[0]
			var operand_one_int uint16
			operator = instruction[1]
			operand_two = instruction[2]
			var operand_two_int uint16

			if _, err := strconv.Atoi(operand_one); err == nil {
				value, _ := strconv.ParseInt(operand_one, 10, 16)
				operand_one_int = uint16(value)

			} else {
				operand_one_int = circuit[operand_one]
			}

			if _, err := strconv.Atoi(operand_two); err == nil {
				value, _ := strconv.ParseInt(operand_two, 10, 16)
				operand_two_int = uint16(value)

			} else {
				operand_two_int = circuit[operand_two]
			}

			if operator == "AND" {
				circuit[destination] = operand_one_int & operand_two_int
			} else if operator == "OR" {
				circuit[destination] = operand_one_int | operand_two_int
			} else if operator == "LSHIFT" {
				circuit[destination] = operand_one_int << operand_two_int
			} else if operator == "RSHIFT" {
				circuit[destination] = operand_one_int >> operand_two_int
			}

		}

		fmt.Println(circuit[destination])
	}

}
