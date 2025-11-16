package main

import (
	"fmt"
	"go-grpc-client-server/server/math"
)

func main() {
	fmt.Println("Hello world from the server!")

	result := math.Add(2, 2)
	fmt.Println("Two plus two is equal to: ", result)
}
