package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func EchoServer() {
	e := echo.New()

	e.GET("/",hello)

	e.Logger.Fatal(e.Start(":8080"))

}

func hello(c echo.Context) error {
	return c.String(200,"hello world")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func Server() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func main() {
	Server()
}