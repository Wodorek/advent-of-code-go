package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	i := 0
	for {
		hash := getHash(input + strconv.Itoa(i))
		if hash[:5] == "00000" {
			break
		}
		i++
	}

	fmt.Println(i)

}

func getHash(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
