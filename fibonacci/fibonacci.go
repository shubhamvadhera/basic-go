package fibonacci

//CalFibonacci fuction calculates Fibonacci Series recursively
func CalFibonacci(num int) int {
	if num <= 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return (CalFibonacci(num-1) + CalFibonacci(num-2))
	}
}
