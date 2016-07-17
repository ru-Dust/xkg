package main

import (
	"fmt"

	"github.com/LuccaPrado/xkg"
)

func main() {
	var keys = make(chan int, 100)

	go xkg.StartXGrabber(keys)

	for {
		keycode := <-keys

		if key, ok := xkg.KeyMap[keycode]; ok {
			fmt.Printf("[%s]", key)
		}
	}
}
