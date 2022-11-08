package main

import "fmt"

func main() {
	strs := []string{"abc", "dfg"}
	fmt.Printf("%s\n", strs)

	nums := []int{1, 2, 3}
	fmt.Printf("%d", nums)

	str := processUserName("广州17：用户刘敏健大阿哥撒大噶是的创建了1台边缘虚拟机")
	fmt.Println(str)

	return
}

func processUserName(message string) string {
	msgRune := []rune(message)
	idx1, idx2 := 0, 0
	for i := 1; i < len(msgRune)-1; i++ {
		if msgRune[i-1] == '用' && msgRune[i] == '户' {
			idx1 = i + 1
		} else if msgRune[i] == '创' && msgRune[i+1] == '建' || msgRune[i] == '购' && msgRune[i+1] == '买' {
			idx2 = i
			break
		}
	}

	var newMsgRune []rune
	newMsgRune = append(newMsgRune, msgRune[:idx1+1]...)
	newMsgRune = append(newMsgRune, []rune("**")...)
	newMsgRune = append(newMsgRune, msgRune[idx2:]...)
	return string(newMsgRune)
}
