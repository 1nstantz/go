package main

import "fmt"

func main() {
	fmt.Printf("hello world")
	var code, endDate = 123, "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var targetUrl = fmt.Sprintf(url, code, endDate)
	fmt.Printf(targetUrl)
}
