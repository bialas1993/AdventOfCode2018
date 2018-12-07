package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const (
	InputFile  = "input"
	MatrixSize = 1000
)

var (
	fabricMatrix     = [][]int{}
	duplicates   int = 0
)

type square struct {
	ID        int
	positionX int
	positionY int
	width     int
	height    int
}

// Parsing coordinates
// example: #6 @ 362,248: 24x10
func NewSquare(s string) *square {
	re := regexp.MustCompile(`#([0-9]+)\s@\s([0-9]+),([0-9]+):\s([0-9]+)x([0-9]+)`)
	match := re.FindStringSubmatch(s)

	id, _ := strconv.Atoi(match[1])
	positionX, _ := strconv.Atoi(match[2])
	positionY, _ := strconv.Atoi(match[3])
	width, _ := strconv.Atoi(match[4])
	height, _ := strconv.Atoi(match[5])

	return &square{id, positionX, positionY, width, height}
}

func main() {
	vector := []int{}

	for j := 0; j < MatrixSize; j++ {
		vector = append(vector, 0)
	}

	for i := 0; i < MatrixSize; i++ {
		v := append(vector[:0:0], vector...)
		fabricMatrix = append(fabricMatrix, v)
	}

	_, fileName, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(fileName), InputFile)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(inpBuff), "\n")
	for _, row := range rows {
		appenSquare(NewSquare(row))
	}

	for i := 0; i < MatrixSize; i++ {
		for j := 0; j < MatrixSize; j++ {
			if fabricMatrix[i][j] > 1 {
				duplicates++
			}
		}
	}

	println("Duplicates: ", duplicates)
}

func appenSquare(s *square) {
	for i := s.positionY; i < s.positionY+s.height; i++ {
		for j := s.positionX; j < s.positionX+s.width; j++ {
			fabricMatrix[i][j]++
		}
	}
}
