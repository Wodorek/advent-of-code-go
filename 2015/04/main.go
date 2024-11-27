package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/wodorek/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func main() {

	util.PrintSolution(1, p1())
	util.PrintSolution(2, p2())

}

func getHash(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func p1() string {
	i := 0
	for {
		hash := getHash(input + strconv.Itoa(i))
		if hash[:5] == "00000" {
			break
		}
		i++
	}

	return fmt.Sprintf("%d", i)
}

func p2() string {
	i := 0
	for {
		hash := getHash(input + strconv.Itoa(i))
		if hash[:6] == "000000" {
			break
		}
		i++
	}

	return fmt.Sprintf("%d", i)
}
