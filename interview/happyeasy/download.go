package main
/*
	1.并发执行tasks中的任务，顺序输出结果
	2.如果任务是 下载任务，输出任务名称
*/
type Task interface {
	Run()(int,error)
}

type DownloadTask struct {
	Name string
}

func (task *DownloadTask)Run()(int,error)  {
	// run...

	return 0,nil
}


func main() {
	// 里边包含有多种任务，
	var tasks []Task

	// 补全代码
}