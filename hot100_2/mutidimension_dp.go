package hot100_2

// 62. 不同路径
// 输入：m = 3, n = 2
// 输出：3
// 解释：
// 从左上角开始，总共有 3 条路径可以到达右下角。
// 1. 向右 -> 向下 -> 向下
// 2. 向下 -> 向下 -> 向右
// 3. 向下 -> 向右 -> 向下
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 初始状态
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 64. 最小路径和
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	// 最小路径
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// 5. 最长回文子串
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 中心扩散法
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		// 偶数
		s1 := extend(i, i, s)
		s2 := extend(i, i+1, s)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}
func extend(l, r int, s string) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

// 动态归纳法
func longestPalindromeDp(s string) string {
	n := len(s)
	// dp[i][j] 子出串i-j是否是回文子串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	start := 0
	maxLen := 0
	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
					start = i
					maxLen = j - i + 1
				}
			}
		}
	}
	return s[start : maxLen+1]
}

// 516. 最长回文子序列,不需要连续
// 输入：s = "bbbab"
// 输出：4
// 解释：一个可能的最长回文子序列为 "bbbb" 。
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}
	for i := 1; i <= len(s); i++ {
		dp[i][i] = 1
	}

	for i := len(s); i > 0; i-- {
		for j := i + 1; j <= len(s); j++ {
			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[1][len(s)]
}

// 5. 最长回文子串
func longestPalindromeV2(s string) string {
	// dp[i][j] 子出串i-j是否是回文子串
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	start := 0
	maxLen := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					if j-i > maxLen {
						maxLen = j - i
						start = i
					}
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					if j-i > maxLen {
						maxLen = j - i
						start = i
					}
				}
			}
		}
	}
	return s[start : start+maxLen+1]
}

// 最长公共子序列
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace" ，它的长度为 3 。
func longestCommonSubsequenceV1(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}
func longestPalindromeV3(s string) string {
	var res string
	extend := func(l, r int) string {
		for {
			if l < 0 || r >= len(s) || s[l] != s[r] {
				break
			}
			l--
			r++
		}
		return s[l+1 : r]
	}
	for i := 0; i < len(s); i++ {
		s1 := extend(i, i)
		s2 := extend(i, i+1)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

// 1143.最长公共子序列
// 输入：text1 = "abcde", text2 = "ace"
// 输出：3
// 解释：最长公共子序列是 "ace"，它的长度为 3。
func longestCommonSubsequenceV2(text1 string, text2 string) int {
	// dp[i][j]：长度为[0, i - 1]的字符串text1与长度为[0, j - 1]的字符串text2的最长公共子序列为dp[i][j]
	n := len(text1)
	m := len(text2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// i-1 为了统计方便，i从1开始。 思考一下，为什么不需要考虑dp[i][j-1] 和dp[i-1][j]的情况
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

// 72. 编辑距离
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
func minDistance(text1 string, text2 string) int {
	// dp[i][j] 表示将 word1 的前 i 个字符转换成 word2 的前 j 个字符所需的最小操作次数
	n := len(text1)
	m := len(text2)
	// 有一个字符串为空串
	if n*m == 0 {
		return n + m
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				// 相同的时候
				dp[i][j] = dp[i-1][j-1]
			} else {
				// dp[i-1][j-1]是替换操作，就是相同的时候+1次替换。 因为替换为相同元素后
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[n][m]
}

// intention
// execution
// 8 答案6
