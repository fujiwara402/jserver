package main

import (
	"./jserver"
	"os"
)

func _main() int {
	err := jserver.Start()
	if err != nil {
		return 1
	}
	return 0
}

func main() {
	os.Exit(_main())
}
