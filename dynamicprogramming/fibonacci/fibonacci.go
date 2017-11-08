package fibonacci

func fibonacci(n int) int {
	options := []func(n int) int{fib_r, fib_c_driver, fib_dp, fib_ultimate}
	return options[3](n)
}

/*---------------------------------------------------------------------------*/
/* Do not need to store all the intermediate values for the entire period of execution.
   Reduces the storage demands to constant space. */
func fib_ultimate(n int) int {
	back2, back1 := 0, 1
	next := 0 // placeholder for sum

	if n == 0 {
		return 0
	}
	for i := 2; i < n; i++ {
		next = back2 + back1
		back2 = back1
		back1 = next
	}
	return back2 + back1
}

/*---------------------------------------------------------------------------*/
/* Do not use recursion. */
func fib_dp(n int) int {
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

/*---------------------------------------------------------------------------*/
/* Still use recursion but look up in the table firstly. */
func fib_c(n int, f []int) int {
	if f[n] == -1 {
		f[n] = fib_c(n-1, f) + fib_c(n-2, f)
	}
	return f[n]
}

func fib_c_driver(n int) int {
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = -1
	}
	return fib_c(n, f)
}

/*---------------------------------------------------------------------------*/
/* Exponential time complexity! */
func fib_r(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib_r(n-1) + fib_r(n-2)
}
