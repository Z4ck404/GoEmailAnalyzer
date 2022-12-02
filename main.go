package main

import (
	cmd "emailanlyzer/m/commands"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Scan cmd.ScanCLI `cmd:"" help:"scan an eml file"`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	// cmds will be called here
	}
}
