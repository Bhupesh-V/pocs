package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
	// Only works when we do a main module build (doesn't work with direct go run)
	// See https://github.com/golang/go/issues/51279
	info, _ := debug.ReadBuildInfo()
	bi, err := debug.ParseBuildInfo(info.String())
	if err != nil {
		log.Fatal(err)
	}
	for _, dep := range bi.Settings {
		fmt.Println(dep.Key, dep.Value)
		// if dep.Key == "vcs.revision" || dep.Key == "vcs.time" {
		// 	fmt.Println(dep.Key, dep.Value)
		// }
	}
}
