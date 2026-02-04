package main

import (
	"fmt"
)

func main() {

	// Data Type Summary
	types := []struct {
		Type string
		Size uintptr
		Min  string
		Max  string
		Zero string
	}{
		{"int8", 1, "-128", "127", "0"},
		{"int16", 2, "-32768", "32767", "0"},
		{"int32", 4, "-2.1B", "2.1B", "0"},
		{"int64", 8, "-9.2Q", "9.2Q", "0"},
		{"uint8", 1, "0", "255", "0"},
		{"uint16", 2, "0", "65535", "0"},
		{"uint32", 4, "0", "4.3B", "0"},
		{"uint64", 8, "0", "18.4Q", "0"},
		{"float32", 4, "1.4e-45", "3.4e38", "0.0"},
		{"float64", 8, "4.9e-324", "1.8e308", "0.0"},
		{"bool", 1, "false", "true", "false"},
		{"string", 16, "", "", `""`},
	}

	fmt.Println("Type      Size(bytes)  Zero Value")
	fmt.Println("---------------------------------")
	for _, t := range types {
		fmt.Printf("%-8s  %-11d  %-10s\n", t.Type, t.Size, t.Zero)
	}
}
