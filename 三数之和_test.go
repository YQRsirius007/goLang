package pos

import (
	"fmt"
	"testing"
)

func Test_ThreeSum(t *testing.T) {
	// nums := []int{-1, 0, 1, 2, -1, -4, -1, -2}
	// ret := [][]int{
	// 	{-1, -1, 2},
	// 	{-1, 0, 1},
	// }
	// s := threeSum(nums)
	// // if ret!=s{
	// 	t.Errorf("want:%v,got%v", ret, s)
	// // }

}

func getBestTimeWindow(usersPerHour []int, threshold int) []int {
	// TODO:请在此补充代码
	ret := []int{-1, -1}
	start := []int{-1}
	end := []int{-1}
	idx, sum := 0, 0
	maxPeriod := 7 * 24

	for i := 0; i < maxPeriod; i++ {
		// if sum+usersPerHour[i]<=threshold{
		// 	sum += usersPerHour[i]
		// 	if start[idx]==-1{
		// 		start[idx]=i
		// 	}
		// 	end[idx] = i
		// }else if usersPerHour[i]<=threshold{
		// 	sum = usersPerHour[i]
		// 	idx++
		// 	start = append(start, i)
		// 	end = append(end, i)
		// 	fmt.Println(i)
		// }else{
		// 	sum = 0
		// 	idx++
		// 	start = append(start, -1)
		// 	end = append(end, -1)
		// }
		// start[idx] = i
		if usersPerHour[i] <= threshold {
			start[idx], end[idx] = i, i
			sum = usersPerHour[i]
			for j := i + 1; j < 2*maxPeriod; j++ {
				if sum+usersPerHour[j%maxPeriod] <= threshold {
					sum += usersPerHour[j%maxPeriod]
					end[idx] = j
				} else {
					break
				}
			}
			fmt.Println(start[idx], end[idx])
			idx++
			start = append(start, -1)
			end = append(end, -1)
		}
	}
	lenMax := 0
	for i := 0; i <= idx; i++ {
		fmt.Println(start[i], end[idx], lenMax, end[i]-start[i])
		if start[i] != -1 && lenMax < end[i]-start[i] && maxPeriod > end[i]-start[i] {
			lenMax = end[i] - start[i]
			ret = []int{start[i], end[i] % maxPeriod}
		}
	}
	return ret
}
