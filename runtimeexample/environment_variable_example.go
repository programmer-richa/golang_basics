// Package runtimeexample demonstrate how environment variables can be fetched in go lang.
package runtimeexample

import (
	"runtime"

	"github.com/programmer-richa/golang_basics/utility"
)

// Module Level Flag is used for indentation purpose.
const level = 1

//Data demonstrated usage of GOARCH, GOOS, GOPATH, and GOROOT environment variables.
//  They influence the building of Go programs (see https://golang.org/cmd/go and https://golang.org/pkg/go/build).
//  GOARCH, GOOS, and GOROOT are recorded at compile time and made available by constants or functions in this package,
//  but they do not influence the execution of the run-time system.
func Data() {
	utility.Printh(level, "Using runtime Package")
	utility.Println(level, "GOARCH: ", runtime.GOARCH)
	utility.Println(level, "GOOS: ", runtime.GOOS)
	// utility.Println(level,"GOPATH: ", runtime.GOPATH)
	// utility.Println(level,"GOROOT : ", runtime.GOROOT)
}
