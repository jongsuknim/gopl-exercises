package main

import (
	"flag"
	"fmt"
	"github/jongsuknim/gopl-exercises/ch02/01/tempconv"
)

type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "℃":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "℉":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

type kelvinFlag struct{ tempconv.Kelvin }

func (f *kelvinFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "K":
		f.Kelvin = tempconv.Kelvin(value)
		return nil
	case "C", "℃":
		f.Kelvin = tempconv.CtoK(tempconv.Celsius(value))
		return nil
	case "F", "℉":
		f.Kelvin = tempconv.CtoK(tempconv.FToC(tempconv.Fahrenheit(value)))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func KelvinFlag(name string, value tempconv.Kelvin, usage string) *tempconv.Kelvin {
	f := kelvinFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Kelvin
}

var c = CelsiusFlag("c", 20.0, "the celsius temperture")
var k = KelvinFlag("k", 20.0, "the kelvin temperture")

func main() {
	flag.Parse()
	fmt.Println(*c)
	fmt.Println(*k)
}
