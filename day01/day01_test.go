package main

import (
	"strings"
	"testing"
)

func TestResults(t *testing.T) {
	cases := map[string][]string{
		"+3": []string{"+1", "+1", "+1"},
		"+0": []string{"+1", "+1", "-2"},
		"-6": []string{"-1", "-2", "-3"},
	}

	for result, seq := range cases {
		globalSequence := NewSequence("+0")
		for _, s := range seq {
			globalSequence = globalSequence.Apply(NewSequence(strings.Trim(s, "")))
		}

		assertEqual(t, result, globalSequence.String())
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
