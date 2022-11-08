/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	var strs []string
	str_arr := strings.Split(path, "/")
	for _, item := range str_arr {
		if item == ".." {
			if len(strs) == 0 {
				continue
			} else {
				strs = strs[:len(strs)-1]
			}
		} else if item == "." || item == "" {
			continue
		} else {
			strs = append(strs, item)
		}
	}
	return "/" + strings.Join(strs, "/")
}

// @lc code=end

