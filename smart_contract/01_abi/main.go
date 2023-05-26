package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 600; i++ {
		x, y := 3, 4
		if x+y > 5 {
			fmt.Println(55555555)
			break
		}
		time.Sleep(time.Second)
		fmt.Println("sleeping...")
	}

}
