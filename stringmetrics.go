package stringmetrics

func makeMatrix(n, m int) [][]int {
  matrix := make([][]int, n)
  for i := 0; i < n; i++ {
    matrix[i] = make([]int, m)
  }
  return matrix
}

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
  n, m := len(a) + 1, len(b) + 1
  d := makeMatrix(n, m)

  for i := 0; i < n; i++ {
    d[i][0] = i
  }
  for i := 0; i < m; i++ {
    d[0][i] = i
  }

  for i := 1; i < n; i++ {
    for j := 1; j < m; j++ {
      if a[i-1] == b[j-1] {
        d[i][j] = d[i-1][j-1]
      } else {
        d[i][j] = minInt(d[i-1][j] + 1, 
          d[i][j-1] + 1, 
          d[i-1][j-1] + 1)
      }
    }
  }

  return d[n-1][m-1]
}

func Levenshtein2(a, b string) int {
  n, m := len(a), len(b)
  if n > m {
    // make sure n <= m, to use the minimum space
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