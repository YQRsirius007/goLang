package pos

import "fmt"

func guilin() {
	// please define the input here.
	// For example: r := bufio.NewReader(os.Stdin) input, err := r.ReadString('\n')
	// please finish the function body here.
	// please define the output here. For example: fmt.Println(input)
	n, a := 0, 0
	fmt.Scanln(&n)
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, n)
		for j := i; j < n; j++ {
			fmt.Scan(&a)
			nums[i][j] = a
		}
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == 0 {
				dp[j] = nums[i][j]
			} else {
				//fmt.Println("dp", dp)
				//fmt.Println(i, j, dp[i] ,nums[i][j])
				dp[j] = min(dp[j], dp[i-1]+nums[i][j])
			}
		}
	}
	fmt.Println(dp[n-1])
}
