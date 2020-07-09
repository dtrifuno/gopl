package main

import (
	"bufio"
	"fmt"
	"gopl/ch2/ex2/conv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			handleNumber(parseNumber(arg))
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			handleNumber(parseNumber(input.Text()))
		}
	}
}

func parseNumber(s string) float64 {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	return t
}

func handleNumber(t float64) {
	temperatureConversions(t)
	weightConversions(t)
	lengthConversions(t)
	fmt.Println()
}

func temperatureConversions(t float64) {
	c := conv.Celsius(t)
	f := conv.Fahrenheit(t)
	fmt.Printf("%s = %s, %s = %s\n", c, c.ToF(), f, f.ToC())
}

func weightConversions(t float64) {
	kg := conv.Kilograms(t)
	lbs := conv.Pounds(t)
	fmt.Printf("%s = %s, %s = %s\n", kg, kg.ToPounds(), lbs, lbs.ToKilograms())
}

func lengthConversions(t float64) {
	m := conv.Meters(t)
	ft := conv.Feet(t)
	fmt.Printf("%s = %s, %s = %s\n", m, m.ToFeet(), ft, ft.ToMeters())
}
