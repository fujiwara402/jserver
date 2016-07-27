package main

import (
	"os"

	"./jserver"
)

// 終了ステータスを返すようにする

func _main()int{
	err := jserver.Start() 
	err != nil {
		return 1
	}
	return 0
}

func main() {
	os.Exit(_main())
}
