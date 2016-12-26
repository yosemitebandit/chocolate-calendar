package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type char struct {
	value byte
}

func main() {
	input := "uqwqemis"
	letters := "abcdef"
	var password [8]char
	i := 0
	for {
		data := fmt.Sprintf("%s%s", input, strconv.Itoa(i))
		hasher := md5.New()
		hasher.Write([]byte(data))
		hash := hex.EncodeToString(hasher.Sum(nil))
		if hash[0:5] == "00000" {

			// Print what we got.
			fmt.Println(hash)
			index := string(hash[5])

			// If the index is a letter, skip.
			if strings.Contains(letters, index) {
				fmt.Println("letter!")
				fmt.Println()
				i += 1
				continue
			}

			// If the index is greater than the pw length, skip.
			integerIndex, _ := strconv.Atoi(index)
			if integerIndex > len(password)-1 {
				fmt.Println("index too large!")
				fmt.Println()
				i += 1
				continue
			}

			// If we've already filled that index, skip.
			if !((char{}) == password[integerIndex]) {
				fmt.Println("already have that index!")
				fmt.Println()
				i += 1
				continue
			}

			var c char
			c.value = hash[6]
			password[integerIndex] = c

			//Check if array has nil.
			nilPresent := false
			for _, value := range password {
				if (char{}) == value {
					nilPresent = true
				}
			}
			if !nilPresent {
				break
			}
			fmt.Println()
		}
		i += 1
	}

	fmt.Println("Part 2 Solution:")
	for _, c := range password {
		fmt.Print(string(c.value))
	}
}
