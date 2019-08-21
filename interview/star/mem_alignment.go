package main

// TODO 多看 内存对齐问题，还未理解

type T1 struct {
	b int32
	a int8
	c int16
}

type T2 struct {
	a int8
	b int32
	c int16
}

/* 分别占用多少内存 */
