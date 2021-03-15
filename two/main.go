package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var count = 0

func main() {
	b, err := ioutil.ReadFile("inputs.txt")
	if err != nil {
		fmt.Print("Error: ", err)
		os.Exit(1)
	}
	passwords := strings.Split(string(b), "\n")

	for _, password := range passwords {

		countAndTarget := strings.Split(password, " ")
		pwd := countAndTarget[2]
		minMax := countAndTarget[0]
		target := rune(countAndTarget[1][0])

		minMaxSplit := strings.Split(minMax, "-")
		min, err := strconv.Atoi(minMaxSplit[0])
		if err != nil {
			fmt.Print(err)
		}
		max, err := strconv.Atoi(minMaxSplit[1])
		if err != nil {
			fmt.Print(err)
		}

		chars := make(map[rune]int)
		for _, char := range pwd {
			chars[char]++
		}
		if chars[target] >= min && chars[target] <= max {
			count++
		}

	}

	fmt.Println(count)
}
