package stringmetrics

import (
  "testing"
  "github.com/sdegutis/go.assert"
)

var LOREM_ONE string = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
var LOREM_TWO string = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est different."

func Test_Levenshtein(t *testing.T) {
  assert.Equals(t, Levenshtein("this is one", "this is two"), 3)
  assert.Equals(t, Levenshtein("this is one", "this is ono"), 1)
}

func Test_Levenshtein_Long(t *testing.T) {
  assert.Equals(t, Levenshtein(LOREM_ONE, LOREM_TWO), 8)
}

func Benchmark_Levenshtein_Short(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Levenshtein("this is one", "this is two")
  }
}

func Benchmark_Levenshtein_Long(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Levenshtein(LOREM_ONE, LOREM_TWO)
  }
}