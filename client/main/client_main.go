package main

import (
	"fmt"

	"go-grpc-client-server/client/math"
)

func main() {
	fmt.Println("Hello world from the client!")

	result := math.Add(3, 3)
	fmt.Println("Two plus two is equal to: ", result)
}
