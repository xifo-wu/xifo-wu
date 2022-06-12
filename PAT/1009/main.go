package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin) // 标准输入输出
	str, _ = reader.ReadString('\n')    // 回车结束
	str = strings.TrimSpace(str)        // 去除最后一个空格

	arr := strings.Split(str, " ")
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}

	fmt.Println(strings.Join(arr, " "))
}
