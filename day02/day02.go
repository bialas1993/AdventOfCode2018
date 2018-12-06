package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
)

const (
	InputFile = "input"
)

func main() {
	check := map[int]int{
		2: 0,
		3: 0,
	}

	_, fileName, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(fileName), InputFile)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	codes := strings.Split(string(inpBuff), "\n")

	for _, s := range codes {
		fmt.Printf("%s:\n", s)
		list := []rune{}
		checker := map[int]bool{
			2: false,
			3: false,
		}

		for _, ch := range []rune(s) {
			if false == inSlice(ch, list) {
				list = append(list, ch)
				num := strings.Count(s, string(ch))
				if _, ok := checker[num]; ok {
					checker[num] = true
				}
			}
		}

		if checker[2] == true {
			check[2] = check[2] + 1
		}

		if checker[3] == true {
			check[3] = check[3] + 1
		}
	}

	println("Result: ", check[2]*check[3])
}

func inSlice(a rune, list []rune) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
