package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	secretKey := strings.TrimSpace(string(data))
	var number int
	for {
		hash := md5.Sum([]byte(secretKey + strconv.Itoa(number)))
		hashString := fmt.Sprintf("%x", hash)

		if strings.HasPrefix(hashString, "000000") { // Checking for six zeroes
			fmt.Printf("%d\n", number)
			break
		}
		number++
	}
}
