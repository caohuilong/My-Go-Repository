package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Byte     uint64 = 1
	Kilobyte        = Byte * 1000
	Megabyte        = Kilobyte * 1000
	Gigabyte        = Megabyte * 1000
	Terabyte        = Gigabyte * 1000
	Petabyte        = Terabyte * 1000
	Exabyte         = Petabyte * 1000
)

func main() {
	fmt.Println(ByteSize(12))
}

func ByteSize(bytes uint64) string {
	unit := ""
	value := float64(bytes)

	switch {
	case bytes >= Exabyte:
		unit = "EB"
		value = value / float64(Exabyte)
	case bytes >= Petabyte:
		unit = "PB"
		value = value / float64(Petabyte)
	case bytes >= Terabyte:
		unit = "TB"
		value = value / float64(Terabyte)
	case bytes >= Gigabyte:
		unit = "GB"
		value = value / float64(Gigabyte)
	case bytes >= Megabyte:
		unit = "MB"
		value = value / float64(Megabyte)
	case bytes >= Kilobyte:
		unit = "KB"
		value = value / float64(Kilobyte)
	case bytes >= Byte:
		unit = "B"
	case bytes == 0:
		return "0"
	}

	result := strconv.FormatFloat(value, 'f', 2, 64)
	result = strings.TrimSuffix(result, ".00")
	return result + unit
}