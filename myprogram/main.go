package main

import (
	"fmt"

	"github.com/AaronDuda/gunner/client"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(client.Hello)
}
