package main

func main() {
	i := GetValue()
	/* 只有interface可以类型断言*/
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}

func GetValue() int {
	return 1
}
