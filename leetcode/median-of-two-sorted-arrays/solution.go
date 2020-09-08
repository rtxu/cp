import "math"

func findMedianSortedArrays(A []int, B []int) float64 {
	m, n := len(A), len(B)
	if m > n {
		A, B, m, n = B, A, n, m
	}

	left, right := 0, m
	for left <= right {
		i := (left + right) / 2
		j := (m+n)/2 - i

		if (i == 0 || j == n || A[i-1] <= B[j]) && (j == 0 || i == m || B[j-1] <= A[i]) {
			if (m+n)%2 == 0 {
				var lmax, rmin float64
				if i == 0 {
					lmax = float64(B[j-1])
				} else if j == 0 {
					lmax = float64(A[i-1])
				} else {
					lmax = math.Max(float64(A[i-1]), float64(B[j-1]))
				}
				if i == m {
					rmin = float64(B[j])
				} else if j == n {
					rmin = float64(A[i])
				} else {
					rmin = math.Min(float64(A[i]), float64(B[j]))
				}
				return (lmax + rmin) / 2
			} else {
				if i == m {
					return float64(B[j])
				} else if j == n {
					return float64(A[i])
				} else {
					return math.Min(float64(A[i]), float64(B[j]))
				}
			}
		} else if !(i == 0 || j == n || A[i-1] <= B[j]) {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return -1
}
