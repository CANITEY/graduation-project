package main

import (
	"fmt"
	"time"
)

func main() {
	createTime := time.Now().String()
	fmt.Printf("createTime: %v\n", createTime)
}
