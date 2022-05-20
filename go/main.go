package main

import (
	"fmt"
	"os"
	c "tdd/stocks"
)

func main() {
	//get command parameter
	input := os.Args[len(os.Args)-1]

	result := c.ExecuteCompressionLossLess(input)
	fmt.Println(result)
}
