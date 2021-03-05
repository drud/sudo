// +build windows

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var mode string
	var spawn bool
	flag.StringVar(&mode, "mode", "", "mode")
	flag.BoolVar(&spawn, "spawn", false, "spawn")
	flag.Parse()

	args := flag.Args()
	if flag.NArg() == 1 && args[1] == "--version" {
		fmt.Printf("sudo.exe version %s", Version)
	}

	if mode != "" {
		os.Exit(client(mode, args))
	}
	if spawn {
		if flag.NArg() == 0 {
			args = []string{"cmd"}
		}
		os.Exit(start(args))
	}
	if flag.NArg() == 0 {
		args = []string{"cmd", "/c", "start"}
	}
	os.Exit(server(args))
}
