package main

import "time"

func main() {
	cfgDemo()
	<-time.After(time.Second * 100)
}
