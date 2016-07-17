xkg - X Keyboard Grabber
========================

[![GoDoc](https://godoc.org/github.com/go-xkg/xkg?status.png)](https://godoc.org/github.com/go-xkg/xkg)

## Installation

    go get github.com/LuccaPrado/xkg

## Usage example:

```go
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
```

## License

The license of the project is [The MIT License (MIT)](https://github.com/henriquemenezes/xkg/blob/master/LICENSE).
