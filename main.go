package main

import "os"

import "github.com/RedCuckoo/UniBDK-go/unindexer/cli"

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
