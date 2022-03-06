package goHelp

import "strings"

func initIntArray(n int) []int {
	result := make([]int, n)
	for i := 0; i <= sizeR-1; i++ {
		result[i] = 0
	}
	return result
}

func initStrArray(n int) []string {
	result := make([]int, n)
	for i := 0; i <= n-1; i++ {
		result[i] = 0
	}
	return result
}

func getKeysIntS(m map[int]string) []int {
	keys := make([]int, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func getKeysIntI(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func initMirror2dStr(arr [][]str) [][]str {
	var arrWidth int = len(arr[0][:]) - 1
	var arrHeight int = len(arr[:])
	result := make([][]str, arrHeight)
	for i := range result {
		result[i] = make([]str, arrWidth+1)
	}
	return result
}

func initMirror2dInt(arr [][]int) [][]int {
	var arrWidth int = len(arr[0][:]) - 1
	var arrHeight int = len(arr[:])
	result := make([][]int, arrHeight)
	for i := range result {
		result[i] = make([]int, arrWidth+1)
	}
	return result
}

func addAtOffset1dInt(src, dst []int, off int) []int {
	sizeS := len(src) - 1
	for i := 0; i <= sizeS; i++ {
		dst[off+i] = src[i]
	}
	return dst
}

func addAtOffset1dStr(src, dst []string, off int) []string {
	sizeS := len(src) - 1
	for i := 0; i <= sizeS; i++ {
		dst[off+i] = src[i]
	}
	return dst
}

func removeQuotesStrSlices(s string) string {
	builder := strings.Builder{}
	quoteEsc := `"`
	replacement := strings.ReplaceAll(s, ", ", "\", \"")
	builder.WriteString(quoteEsc + replacement + quoteEsc)
	return builder.String()
}

func evalSize2dInt(arr [][]int) (int, int) {
	sizeX := len(arr) - 1
	sizeY := len(arr[:]) - 1
	return sizeY, sizeX
}

func evalSize2dStr(arr [][]string) (int, int) {
	sizeX := len(arr) - 1
	sizeY := len(arr[:]) - 1
	return sizeY, sizeX
}

func largestInt(arr []int) int {
	var n int = 0
	for _, m := range arr {
		if m > n {
			n = m
		}
	}
	return n
}

func largestStr(arr []string) int {
	var n int = 0
	for _, m := range arr {
		if len(m) > n {
			n = m
		}
	}
	return n
}
