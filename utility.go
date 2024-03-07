package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mattn/go-isatty"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	white  = "\033[97m"
)

// Helper function that can be called with a variable to make it appear that the variable is used.
func maybeUnused(x ...interface{}) {}

// Returns a string that consists of [text] followed by zero or more instances of [pad] such that the total length of the returned string is at least [width] characters.
func padStringToWidth(text string, width int, padRune rune) string {
	dots := max(0, width-len(text))
	return text + strings.Repeat(string(padRune), dots)
}

// Scans the file and/or folder paths passed in [paths], applies filters to the files and folders found therein, and returns a slice containing the fully-qualified paths.
func scanAndFilterPaths(paths []string, includeFolder func(basename, fullPath string) bool, includeFile func(basename, fullPath string) bool) []string {
	projectPaths := []string{}
	stack := append([]string{}, paths...)

	for len(stack) > 0 {
		inputPath := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fileInfo, err := os.Stat(inputPath)
		if err != nil {
			os.Stderr.WriteString(fmt.Errorf("can't stat %s: %w", inputPath, err).Error())
			continue
		}

		if !fileInfo.IsDir() {
			if includeFile(filepath.Base(inputPath), inputPath) {
				projectPaths = append(projectPaths, inputPath)
			}
			continue
		}

		// It's a folder.
		entries, err := os.ReadDir(inputPath)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			entryName := entry.Name()
			fullPath := filepath.Join(inputPath, entryName)
			if entry.IsDir() {
				if includeFolder(entryName, fullPath) {
					stack = append(stack, fullPath)
				}
			} else if includeFile(entryName, fullPath) {
				projectPaths = append(projectPaths, fullPath)
			}
		}

	}

	return projectPaths
}

// Calls a function [fn] to iterate over a map whose keys have been sorted case-insensitively.
func iterateOverCISortedMap[V any](m map[string]V, fn func(string, V)) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})
	for _, k := range keys {
		fn(k, m[k])
	}
}

// Iterates over a map and calls the provided function for each key-value pair.
func iterateOverMap[V any](m map[string]V, fn func(string, V)) {
	for k, v := range m {
		fn(k, v)
	}
}

// Takes a byte slice and prints its contents in a hexadecimal format to the standard output.
// It iterates through the byte slice, formatting the bytes and corresponding characters, and then prints them with color formatting.
// The function uses color formatting for better visualization of the output.a hexadecimal format to stdout.
func dumpHexImpl(data []byte, offsetColour, bytesColour, charsColour, resetColour string) {
	const columns = 32
	byteCount := len(data)
	for location := 0; location < byteCount; {
		bytes := []byte{}
		chars := []byte{}
		rowColumns := min(columns, byteCount-location)
		for n := 0; n < rowColumns; n++ {
			value := data[location+n]
			bytes = append(bytes, fmt.Sprintf("%02x ", value)...)
			if value >= 32 && value <= 127 {
				chars = append(chars, value)
			} else {
				chars = append(chars, '.')
			}
		}
		fmt.Printf(offsetColour+"%08x "+bytesColour+"%s%*s"+charsColour+"%s"+resetColour+"\n", location, bytes, (columns-rowColumns)*3, "", chars)
		location += rowColumns
	}
}

func dumpHex(data []byte) {
	isAtty := isatty.IsTerminal(os.Stdout.Fd())
	if isAtty {
		dumpHexImpl(data, green, cyan, yellow, reset)
	} else {
		dumpHexImpl(data, "", "", "", "")
	}
}

// De-duplicate a sorted slice of values of type T.
func dedup[T comparable](tracks []T) []T {
	if len(tracks) == 0 {
		return tracks
	}

	copy := append([]T{}, tracks...)

	i, j := 0, 1
	for j < len(copy) {
		if copy[i] == copy[j] {
			j++
		} else {
			i++
			copy[i] = copy[j]
			j++
		}
	}
	return copy[:i+1]
}

// Calculates the maximum length (in bytes) of the string keys in a map.
func calculateMaximumKeyWidth[T any](m map[string]T) int {
	maximumKeyWidth := 0
	for key := range m {
		keyLength := len(key)
		if keyLength > maximumKeyWidth {
			maximumKeyWidth = keyLength
		}
	}
	return maximumKeyWidth
}
