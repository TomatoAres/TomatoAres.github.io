package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s:= strings.NewReader("helllo")
	br := bufio.NewReader(s)
	fmt.Println(br.Buffered())//0
	br.Peek(1)
	fmt.Println(br.Buffered())//5 or 6

}

/*
TODO 暂时无法理解

// bufio.go 100 ~ 104
n, err := b.rd.Read(b.buf[b.w:])
if n < 0 {
    panic(errNegativeRead)
}
b.w += n
*/
