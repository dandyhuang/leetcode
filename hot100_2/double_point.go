package hot100_2

import (
	"sort"
)

// 283. 移动零
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 88. 合并两个有序数组
// 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出：[1,2,2,3,5,6]
// 解释：需要合并 [1,2,3] 和 [2,5,6] 。
// 合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
func mergeArray(nums1 []int, m int, nums2 []int, n int) {
	r1, r2 := m-1, n-1
	size := m + n - 1
	for r1 >= 0 && r2 >= 0 {
		if nums1[r1] > nums2[r2] {
			nums1[size] = nums1[r1]
			r1--
		} else {
			nums1[size] = nums2[r2]
			r2--
		}
		size--
	}
	for r1 >= 0 {
		nums1[size] = nums1[r1]
		r1--
		size--
	}
	for r2 >= 0 {
		nums1[size] = nums2[r2]
		r2--
		size--
	}
}

// 11. 盛最多水的容器
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49 7*7
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
func maxArea(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	for l < r {
		w := r - l
		h := min(height[l], height[r])
		res = max(res, w*h)
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return res
}

// 15. 三数之和
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	// l== i+1所以-2
	for i := 0; i < len(nums)-2; i++ {
		// 因为单调递增，相等说明之前验证过,不重复的三元组
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				arr := []int{nums[i], nums[l], nums[r]}
				res = append(res, arr)
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	return res
}

// 接雨水
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
// 对于下标 i下雨后水能到达的最大高度等于下标 i两边的最大高度的最小值，
// 下标 i处能接的雨水量等于下标 i处的水能到达的最大高度减去 height[i]。
// 纵向接雨水
func trap(height []int) int {
	n := len(height)
	lHeight := make([]int, n)
	rHeight := make([]int, n)
	lHeight[0] = height[0]
	rHeight[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		lHeight[i] = max(lHeight[i-1], height[i])
	}
	for j := n - 2; j >= 0; j-- {
		rHeight[j] = max(rHeight[j+1], height[j])
	}
	var res int
	for i := 0; i < n; i++ {
		res += min(lHeight[i], rHeight[i]) - height[i]
	}
	return res
}

// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
func trapV2(height []int) int {
	if len(height) < 3 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	result := 0

	for left < right {
		// 更新左右两边的最大高度
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		// 根据木桶效应，取较小的最大高度，计算当前位置的储水量
		if leftMax < rightMax {
			result += leftMax - height[left]
			left++
		} else {
			result += rightMax - height[right]
			right--
		}
	}

	return result
}

// 横向接雨水
// 输入：height = [5,4,2,6,1,0]
// 输出：6
// 解释：上面是由数组 [6] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
// 找上一个更大元素，在找的过程中填坑，灵神出版https://www.bilibili.com/video/BV1VN411J7S7/?vd_source=601fac537a0633b7cf3313f0a02ed170
func trapV3(height []int) int {
	var res = 0
	var stack []int
	for i, h := range height {
		// 相同的时候，也可以进入直接删除就可以了。
		for len(stack) > 0 && height[stack[len(stack)-1]] <= height[i] {
			bottomH := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			wide := i - stack[len(stack)-1] - 1
			leftHeight := min(h, height[stack[len(stack)-1]]) - bottomH
			res += leftHeight * wide
		}
		stack = append(stack, i)
	}
	return res
}
