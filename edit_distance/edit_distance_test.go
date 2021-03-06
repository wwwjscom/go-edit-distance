package edit_distance

import (
	"testing"
)

var testCases = []struct {
	source string
	target string
	distance int
}{
	{"", "a", 1},
	{"a", "aa", 1},
	{"a", "aaa", 2},
	{"", "", 0},
	{"a", "b", 2},
	{"aaa", "aba", 2},
	{"aaa", "ab", 3},
	{"a", "a", 0},
	{"ab", "ab", 0},
	{"a", "", 1},
	{"aa", "a", 1},
	{"aaa", "a", 2},
	{"a.aa$ ", "a", 5},
}

func TestLevenshtein(t *testing.T) {
	for _, testCase := range testCases {
		distance := DistanceForStrings(
				[]rune(testCase.source),
				[]rune(testCase.target),
				DefaultOptions)
		if distance != testCase.distance {
			t.Log(
					"Distance between",
					testCase.source,
					"and",
					testCase.target,
					"computed as",
					distance,
					", should be",
					testCase.distance)
			t.Fail()
		}
	}
}
