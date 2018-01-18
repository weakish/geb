/*
Package geb cross compiles go executables for different os and architectures.

Some combinations of operating system and architecture,
although supported by Go compiler, will be omitted, e.g. darwin on 386.

Usage

Put the geb in $PATH, then run `geb` under your Go project directory.
 */

package main

import (
	"os"
	"os/exec"
	"path"
)


var osArchs = []struct {
	goOs   string
	goArch string
}{
	// windows
	{"windows", "386"},
	{"windows", "amd64"},
	// osx
	{"darwin", "amd64"},
	// bsd
	{"dragonfly", "amd64"},
	{"freebsd", "386"},
	{"freebsd", "amd64"},
	{"freebsd", "arm"},
	{"netbsd", "386"},
	{"netbsd", "amd64"},
	{"netbsd", "arm"},
	{"openbsd", "386"},
	{"openbsd", "amd64"},
	{"openbsd", "arm"},
	// linux
	{"linux", "386"},
	{"linux", "amd64"},
	{"linux", "arm"},
	{"linux", "arm64"},
	{"linux", "ppc64"},
	{"linux", "ppc64le"},
	{"linux", "mips"},
	{"linux", "mipsle"},
	{"linux", "mips64"},
	{"linux", "mips64le"},
	// solaris
	{"solaris", "amd64"},
	// plan9
	{"plan9", "386"},
	{"plan9", "amd64"},
}

func buildExecutable(name string, goOs string, goArch string) {
	var goOsVar string = "GOOS=" + goOs
	var goArchVar string = "GOARCH=" + goArch
	var output string = name + "-" + goOs + "-" + goArch
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
