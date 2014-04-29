package stringmetrics

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Levenshtein(a, b string) int {
	n, m := len(a), len(b)
	if n > m {
		a, b = b, a
		n, m = m, n
	}

	current := make([]int, m+1)
	previous := make([]int, m+1)
	var i, j, add, delete, change int

	for i = 1; i <= m; i++ {
		copy(previous, current)
		for j = 0; j <= m; j++ {
			current[j] = 0
		}
		current[0] = i
		for j = 1; j <= n; j++ {
			if a[j-1] == b[i-1] {
				current[j] = previous[j-1]
			} else {
				add = previous[j] + 1
				delete = current[j-1] + 1
				change = previous[j-1] + 1
				current[j] = min(min(add, delete), change)
			}
		}
	}

	return current[n]
}

func Jaro(a, b string) float64 {
	n, m := len(a), len(b)
	if n == 0 || m == 0 {
		if n == 0 && m == 0 {
			return 1.0
		}
		return 0.0
	}

	if n > m {
		a, b = b, a
		n, m = m, n
	}

	search_range := (n / 2) - 1
	if search_range < 0 {
		search_range = 0
	}

	idx1 := make([]bool, n)
	idx2 := make([]bool, m)

	match := 0
	for i := 0; i < n; i++ {
		start := max(0, i-search_range)
		end := min(i+search_range+1, m)
		for j := start; j < end; j++ {
			if !idx2[j] && a[i] == b[j] {
				idx1[i] = true
				idx2[j] = true
				match++
				break
			}
		}
	}
	if match == 0 {
		return 0.0
	}

	var i, j, trans int
	for i, j, trans = 0, 0, 0; i < n; i++ {
		if idx1[i] {
			for !idx2[j] {
				j++
			}
			if a[i] != b[j] {
				trans++
			}
			j++
		}
	}

	matchf := float64(match)
	return (matchf/float64(n) + matchf/float64(m) + float64(match-trans/2)/matchf) / 3.0
}

func JaroWinkler(a, b string) float64 {
	weight := Jaro(a, b)
	if weight > 0.7 {
		i, j := 0, min(min(len(a), len(b)), 4)
		for ; i < j; i++ {
			if a[i] != b[i] {
				break
			}
		}
		weight += float64(i) * 0.1 * (1.0 - weight)
	}

	return weight
}

func Hamming(a, b string) int {
	if len(a) != len(b) {
		return -1
	}
	sum := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			sum += 1
		}
	}
	return sum
}
