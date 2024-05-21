package main

import "fmt"
import "crypto/md5"
import "strconv"
import "encoding/hex"

func main() {
	puzzle := "REDACTED"
	current_num := 0

	for {

		current_num_string := strconv.Itoa(current_num)

		hash := md5.Sum([]byte(puzzle + current_num_string))

		if hex.EncodeToString(hash[:])[0:6] == "000000" {
			fmt.Println(current_num)
			break
		} else {
			current_num += 1
		}
	}
}
