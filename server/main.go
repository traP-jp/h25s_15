package main

import "time"

func main() {
	for {
		<-time.Tick(time.Second * 3)
	}
}
