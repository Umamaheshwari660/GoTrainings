package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== MIN/MAX VALUES ===")

	// Integer limits
	fmt.Printf("int8:   min=%d, max=%d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16:  min=%d, max=%d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32:  min=%d, max=%d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64:  min=%d, max=%d\n", math.MinInt64, math.MaxInt64)

	fmt.Printf("\nuint8:  max=%d\n", math.MaxUint8)
	fmt.Printf("uint16: max=%d\n", math.MaxUint16)
	fmt.Printf("uint32: max=%d\n", math.MaxUint32)
	fmt.Printf("uint64: max=%d\n", uint64(math.MaxUint64))

	// Float limits
	fmt.Printf("\nfloat32: min=%.5e, max=%.5e\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Printf("float64: min=%.5e, max=%.5e\n", math.SmallestNonzeroFloat64, math.MaxFloat64)

	// Practical demonstration
	var maxInt8 int8 = math.MaxInt8
	var minInt8 int8 = math.MinInt8

	fmt.Printf("\nPractical example:\n")
	fmt.Printf("maxInt8 + 1 = %d (overflow!)\n", maxInt8+1)
	fmt.Printf("minInt8 - 1 = %d (underflow!)\n", minInt8-1)
}
