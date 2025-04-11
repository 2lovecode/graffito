package top100

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	for mask := 0; mask < 1<<uint(len(nums)); mask++ {
		set := []int{}
		for i := 0; i < len(nums); i++ {
			if mask>>uint(i)&1 == 1 {
				set = append(set, nums[i])
			}
		}
		ans = append(ans, set)
	}
	return ans
}

func subsets2(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]

		dfs(cur + 1)
	}
	dfs(0)
	return
}
