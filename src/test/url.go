package main

import (
	"fmt"
	_ "net/url"
)

var urls []string = []string{
	"https://golang.org/src/net/url/url.go?s=13404:13442#L497",
	"http://image16.poco.cn/mypoco/myphoto/20141113/16/174956375201411131659521617469612410_003.jpg?420x631_120",
	"http://p2.pstatp.com/large/2846/8413214320",
	"http://file28.mafengwo.net/M00/A5/87/wKgB6lQP74uAEBvSAAKbsVAdik426.groupinfo.w600.jpeg",
	"http://house.timedg.com/archive.php?aid=18795"}

var escape_url []string = []string{"a", "b", "c"}

func main() {
	var a []string
	//escape_url = make([]string, 10, 10)
	for _, v := range urls {
		//fmt.Printf("The URL is:%s\n", v)
		a = append(escape_url, v)
		//escape_url[i] = url.QueryEscape(v)
		fmt.Println(a)

	}

	/*
		fmt.Println(len(escape_url))
		fmt.Println(cap(escape_url))
		fmt.Println(escape_url)
	*/
}
