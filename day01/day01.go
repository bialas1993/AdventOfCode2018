package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strconv"
	"strings"
)

const (
	InputFile     = "input1"
	OperatorMinus = "-"
	OperatorPlus  = "+"
)

type sequence struct {
	operator string
	position int
}

func NewSequence(seq string) *sequence {
	s := &sequence{}
	bufRune := []rune(strings.Trim(seq, ""))
	s.operator = string(bufRune[0])
	pos, _ := strconv.Atoi(string(bufRune[1:]))
	s.position = pos

	return s
}

func (s *sequence) Apply(ss *sequence) *sequence {
	in := s.position
	if s.operator == OperatorMinus {
		in = in * -1
	}

	out := ss.position
	if ss.operator == OperatorMinus {
		out = out * -1
	}

	sum := out + in
	if sum < 0 {
		return NewSequence(fmt.Sprintf("%s%d", OperatorMinus, sum*-1))
	}

	return NewSequence(fmt.Sprintf("%s%d", OperatorPlus, sum))
}

func (s *sequence) String() string {
	if s.position < 0 {
		return fmt.Sprintf("%s%d", s.operator, s.position*-1)
	}

	return fmt.Sprintf("%s%d", s.operator, s.position)
}

func main() {
	_, fileName, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(fileName), InputFile)
	inpBuff, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	sequences := strings.Split(string(inpBuff), "\r")

	// sequences = []string{"+1", "+1", "+1"}
	// sequences = []string{"+1", "+1", "-2"}
	// sequences = []string{"-1", "-2", "-3"}

	globalSequence := NewSequence("+0")

	for _, seq := range sequences {
		globalSequence = globalSequence.Apply(NewSequence(strings.Trim(seq, "")))
	}

	fmt.Printf("%s\n", globalSequence.String())
}
