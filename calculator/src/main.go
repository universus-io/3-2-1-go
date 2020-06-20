package main

import (
	"io/universus/calculator"
)

func main() {

	server := calculator.NewServer(7001)

	latch := server.Start()

	latch.Wait()

}
