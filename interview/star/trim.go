package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "abaaxy中文cdbbab"
	fmt.Println(s)
	fmt.Println(strings.TrimPrefix(s, "ab"))
	fmt.Println(strings.TrimSuffix(s, "ab"))
	fmt.Println(strings.TrimLeft(s, "ab"))
	fmt.Println(strings.TrimRight(s, "ab"))
	fmt.Println(strings.Trim(s, "ab"))
}

/*
TrimPrefix: 去除前缀ab
TrimSuffix: 去除后缀ab
Trim: 去除最左边和最右边的a或b
TrimLeft: 去除最左边的a或b
TrimRight: 去除最右边的a或b
*/