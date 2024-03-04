package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"slices"
	"strings"

	"github.com/MrSplidge/go-coutil"
)

const (
	alsExtension = ".als"
	cprExtension = ".cpr"
)

type stringFlags []string

// Format the accumulated flags as a string.
func (sf *stringFlags) String() string {
	return fmt.Sprintf("%v", *sf)
}

// Accumulate semicolon separated values into a slice.
func (sf *stringFlags) Set(value string) error {
	*sf = append(*sf, strings.Split(value, ";")...)
	return nil
}

// Main entry point for the program.
func main() {
	var numThreadsFlag = flag.Int("num-threads", runtime.NumCPU(), "The number of worker threads to use.")

	var foldersToIgnore stringFlags
	flag.Var(&foldersToIgnore, "ignore-folders", "A semicolon-separated list of folders to ignore when traversing the hierarchy.")

	var extensions stringFlags
	flag.Var(&extensions, "extensions", "A semicolon-separated list of project file extensions to include when traversing the hierarchy (default .als;.cpr).")

	var cpuProfile string
	flag.StringVar(&cpuProfile, "cpu-prof", "", "Write CPU profile to a file (default.pgo recommended).")

	flag.Parse()

	if len(extensions) == 0 {
		extensions = []string{alsExtension, cprExtension}
	}

	fmt.Printf("Using %d threads.\n", *numThreadsFlag)
	fmt.Println("Ignoring these folders:", foldersToIgnore)
	fmt.Println("Scanning these items:", flag.Args())
	fmt.Println()

	if cpuProfile != "" {
		file, err := os.Create(cpuProfile)
		if err != nil {
			log.Fatal(err)
		}

		if err = pprof.StartCPUProfile(file); err != nil {
			log.Fatal(err)
		}

		defer pprof.StopCPUProfile()
	}

	coutil.WorkPool(
		*numThreadsFlag,
		// Work items to process.
		scanAndFilterPaths(flag.Args(),
			// Exclude certain folders.
			func(basename, fullPath string) bool {
				return !slices.Contains(foldersToIgnore, basename)
			},
			// Include files with certain file extensions.
			func(basename, fullPath string) bool {
				return slices.Contains(extensions, filepath.Ext(basename))
			}),
		// Work item processor.
		func(path string) *projectInformation {
			switch filepath.Ext(path) {
			case alsExtension:
				return examineALS(path)
			case cprExtension:
				return examineCPR(path)
			}
			return nil
		},
		// Results processor.
		func(pi *projectInformation) {
			if pi != nil {
				fmt.Print(pi.String())
			}
		})
}
