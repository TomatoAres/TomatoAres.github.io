package main

import "io"

type MyWriter struct{}

func (m *MyWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

var _ io.Writer = (*MyWriter)(nil)

//请问，var _ io.Writer = (*MyWriter)(nil) 有什么用？

/* 类型转换
和 float32(2) 类似
*/