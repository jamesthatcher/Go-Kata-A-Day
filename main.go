package main

import (
	"code_wars/rainfall"
	"fmt"
	"strconv"
)

var data = "London:Jan 48.0,Feb 38.0,Mar 39.2,Apr 42.2 \n Beijing:Jan 48.0,Feb 38.9,Mar 39.9,Apr 42.2"

func main() {
	r := rainfall.Mean("London", data)
	fmt.Printf(strconv.Itoa(int(r)))
	fmt.Sprintf("OUT: %d", r)
	return
}
