package main

import "fmt"

type People interface {
	Speck(string) string
}

type Teacher struct {}

func (t *Teacher)Speck(talk string) (ret string)  {
	if talk == "you are a bitch."{
		ret = "nihao"
		return
	}else {
		ret = talk
		return
	}
}

func main() {

	var p People = Teacher{}
	//var p People

	talk := "jks"
	fmt.Println(p.Speck(talk))
}
