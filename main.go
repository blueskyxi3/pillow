package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Tick(time.Second)
	for {
		select {
		case <-t:
			fmt.Printf("\r%s----------", time.Now())
		}
	}
}
