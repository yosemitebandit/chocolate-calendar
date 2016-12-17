package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	input := "uqwqemis"
	var password []byte
	i := 0
	for {
		data := fmt.Sprintf("%s%s", input, strconv.Itoa(i))
		hasher := md5.New()
		hasher.Write([]byte(data))
		hash := hex.EncodeToString(hasher.Sum(nil))
		if hash[0:5] == "00000" {
			fmt.Println(i, "->", hash)
			password = append(password, hash[5])
			if len(password) == 8 {
				break
			}
		}
		i += 1
	}

	fmt.Println("Part 1 Solution:", string(password))
}
