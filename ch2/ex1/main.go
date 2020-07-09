package main

import (
	"fmt"
	tc "gopl/ch2/ex1/tempconv"
)

func main() {
	cels := tc.Celsius(30)
	fmt.Printf("%v is %v\n", cels, cels.ToK())

	fahr := tc.Fahrenheit(96)
	fmt.Printf("%v is %v\n", fahr, fahr.ToK())

	kelv := tc.Kelvin(300)
	fmt.Printf("%v is %v, which is %v\n", kelv, kelv.ToC(), kelv.ToF())
}
