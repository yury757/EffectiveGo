package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	fmt.Println(s)
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// 1.17及1.17之前不支持泛型，需要使用1.18及更高版本才能使用
// func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func TestMethod() {
	var a = map[string]int64 {"a": 1, "b": 2, "c": 3}
	fmt.Printf("SumInts result is %d\n", SumInts(a))

	var b = map[string]float64 {"a": 1.0, "b": 2.0, "c": 3.0}
	fmt.Printf("SumFloats result is %f\n", SumFloats(b))
}

func TestGenericMethod() {
        var a = map[string]int64 {"a": 1, "b": 2, "c": 3}
        fmt.Printf("SumIntsOrFloats(a) result is %d\n", SumIntsOrFloats(a))

        var b = map[string]float64 {"a": 1.0, "b": 2.0, "c": 3.0}
        fmt.Printf("SumIntsOrFloats(b) result is %f\n", SumIntsOrFloats(b))
}

func main() {
	TestMethod()
	TestGenericMethod()
}
