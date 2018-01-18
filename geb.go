/*
Package geb cross compiles go executables for different os and architectures.

Currently only supports 4 oses on amd64:

1. darwin
2. freebsd
3. linux
4. windows

Usage

Put the geb in $PATH, then run `geb` under your Go project directory.
 */

package main

import (
	"os"
	"os/exec"
	"path"
)


var osArchs = []struct{
	goOs   string
	goArch string
}{
	// windows
	{"windows", "amd64"},
	// osx
	{"darwin", "amd64"},
	// freebsd
	{"freebsd", "amd64"},
	// linux
	{"linux", "amd64"},
}

func buildExecutable(name string, goOs string, goArch string) {
	var goOsVar string = "GOOS=" + goOs
	var goArchVar string = "GOARCH=" + goArch
	var output string = name + "-" + goOs
	buildCommand := exec.Command("env", goOsVar, goArchVar, "go", "build", "-o", output)
	err := buildCommand.Run()
	if err != nil {
		errorMessage := "An error occurred during compiling " + name + "on " + goOs + "/" + goArch
		os.Stderr.WriteString(errorMessage)
	}
}

func main() {
	var currentWorkingDirectory string
	currentWorkingDirectory, _ = os.Getwd()
	var name string = path.Base(currentWorkingDirectory)
	for _, osArch := range osArchs {
		buildExecutable(name, osArch.goOs, osArch.goArch)
	}
}
