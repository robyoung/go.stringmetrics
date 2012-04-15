package stringmetrics

func minInt(a, b, c int) int {
  min := a
  if b < min {
    min = b
  }
  if c < min {
    min = c
  }
  return min
}

func Levenshtein(a, b string) int {
  n, m := len(a), len(b)
  if n > m {
    a, b = b, a
    n, m = m, n
  }

  current  := make([]int, m+1)
  previous := make([]int, m+1)
  var i, j, add, delete, change int

  for i = 1; i <= m; i++ {
    copy(previous, current)
    for j = 0; j <= m; j++ { current[j] = 0 }
    current[0] = i
    for j = 1; j <= n; j++ {
      if a[j-1] == b[i-1] {
        current[j] = previous[j-1]
      } else {
        add    = previous[j] + 1
        delete = current[j-1] + 1
        change = previous[j-1] + 1
        current[j] = minInt(add, delete, change)
      }
    }
  }

  return current[n]
}