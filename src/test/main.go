package main

import (
	"fmt"
	_ "github.com/jonsen/gotld"
)

var sum int = 30

func main() {
	num := []int{1, 3, 5, 7, 9, 11, 13, 15}
	for i, _ := range num {
		for ii, _ := range num {
			for iii, _ := range num {
				fmt.Printf("%d+%d+%d\n", num[i], num[ii], num[iii])
				if num[i]+num[ii]+num[iii] == sum {

					fmt.Print("OK")
				}
			}
		}
	}
}
