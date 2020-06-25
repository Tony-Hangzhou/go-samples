package main

import (
	"go-samples/pkg/put"
	_ "go-samples/pkg/put"
	_ "go-samples/pkg/watch"
)

func main() {

	put.PutAndGet()

}
