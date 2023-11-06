package leetcode_hot100

// 560. 和为 K 的子数组 前缀和哈希表
func subarraySum(nums []int, k int) int {
	count := 0
	preSum := make([]int, len(nums)+1)
	preSum[0] = nums[0] // preSum[0] 应该为0
	// 通过暴力请求，会发现有问题
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if preSum[i]-preSum[j] == k {
				count++
			}
		}
	}
	return count
}

// 560. 和为 K 的子数组 前缀和哈希表
func subarraySum2(nums []int, k int) int {
	count := 0
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		num := sum - k
		if v, ok := m[num]; ok {
			count += v
		}
		m[sum]++
	}
	return count
}

type MyQueue struct {
	q []int
}

func (q *MyQueue) Push(val int) {
	for len(q.q) > 0 && val > q.q[len(q.q)-1] {
		q.q = q.q[:len(q.q)-1]
	}
	q.q = append(q.q, val)
}

func (q *MyQueue) Pop(val int) {
	if len(q.q) > 0 && q.q[0] == val {
		q.q = q.q[1:]
	}
}

// 239. 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	q := &MyQueue{}
	res := make([]int, 0)
	for i := 0; i < k && i < len(nums); i++ {
		q.Push(nums[i])
	}
	res = append(res, q.q[0])
	for i := k; i < len(nums); i++ {
		q.Pop(nums[i-k])
		q.Push(nums[i])
		res = append(res, q.q[0])
	}
	return res
}

// 76. 最小覆盖子串
func minWindow(s string, t string) string {
	sM := make(map[byte]int)
	tM := make(map[byte]int)

	for i, _ := range t {
		tM[t[i]]++
	}

	count := 0
	var res string
	for i, j := 0, 0; i < len(s); i++ {
		sM[s[i]]++
		if sM[s[i]] <= tM[s[i]] {
			count++
		}
		for j < len(s) && sM[s[j]] > tM[s[j]] {
			sM[s[j]]--
			j++
		}
		if count == len(t) {
			if len(res) == 0 || i-j+1 < len(res) {
				res = s[j : i+1]
			}
		}
	}
	return res
}