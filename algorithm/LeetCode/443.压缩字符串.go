/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

// @lc code=start
func compress(chars []byte) int {
	ans := 0
	begin := 0
	for i := 0; i < len(chars); i++ {
		if i == len(chars)-1 || chars[i] != chars[i+1] {
			chars[ans] = chars[begin]
			ans++
			count := i-begin+1
			if count > 1 {
				str := strconv.Itoa(count)
				for _, c := range str {
					chars[ans] = byte(c)
					ans++
				}
			}
			begin = i+1
		}
	}
	return ans
}

// @lc code=end

