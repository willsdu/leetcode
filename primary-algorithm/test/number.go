package test

func countPrimes(n int) int {
	a := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if a[i] {
			continue
		}
		for j := i + i; j < n; j += i {
			a[j] = true
		}
		count++
	}
	return count
}

func consecutiveNumbersSum(N int) int {
	c := 0
	for i := 0; i*(i+1) < 2*N; i++ {
		if (N*2-i-i*i)%(2*i+2) == 0 {
			c++
		}
	}
	return c
}
