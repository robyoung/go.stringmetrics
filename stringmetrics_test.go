package stringmetrics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var LOREM_ONE string = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
var LOREM_TWO string = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est different."

func Test_Levenshtein(t *testing.T) {
	assert.Equal(t, Levenshtein("this is one", "this is two"), 3)
	assert.Equal(t, Levenshtein("this is one", "this is ono"), 1)
}

func Test_Levenshtein_Long(t *testing.T) {
	assert.Equal(t, Levenshtein(LOREM_ONE, LOREM_TWO), 8)
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

func Test_Jaro(t *testing.T) {
	assert.Equal(t, Jaro("", ""), 1.0)
	assert.Equal(t, Jaro("Brian", ""), 0.0)
	assert.Equal(t, Jaro("Brian", "Brian"), 1.0)
	assert.Equal(t, Jaro("Brian", "Jesus"), 0.0)
	assert.Equal(t, Jaro("Thorkel", "Thorgier"), 0.77976190476190477)
	assert.Equal(t, Jaro("Dinsdale", "D"), 0.70833333333333334)
}

func Test_Jaro_Long(t *testing.T) {
	assert.Equal(t, Jaro(LOREM_ONE, LOREM_TWO), 0.9895633141148835)
}

func Benchmark_Jaro_Short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Jaro("Thorkel", "Thorgier")
	}
}

func Test_JaroWinkler(t *testing.T) {
	assert.Equal(t, JaroWinkler("", ""), 1.0)
	assert.Equal(t, JaroWinkler("Brian", ""), 0.0)
	assert.Equal(t, JaroWinkler("Brian", "Brian"), 1.0)
	assert.Equal(t, JaroWinkler("Brian", "Jesus"), 0.0)
	// slightly boosted higher values
	assert.Equal(t, JaroWinkler("Thorkel", "Thorgier"), 0.8678571428571429)
	assert.Equal(t, JaroWinkler("Dinsdale", "D"), 0.7375)

}

func Test_JaroWinkler_Long(t *testing.T) {
	assert.Equal(t, JaroWinkler(LOREM_ONE, LOREM_TWO), 0.9937379884689301)
}

func Benchmark_JaroWinkler_Short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JaroWinkler("Thorkel", "Thorgier")
	}
}

func Test_Hamming(t *testing.T) {
	assert.Equal(t, Hamming("", ""), 0)
	assert.Equal(t, Hamming("a", ""), -1) // testing strings of unequal length is undefined
	assert.Equal(t, Hamming("Brian", "Jesus"), 5)
}
