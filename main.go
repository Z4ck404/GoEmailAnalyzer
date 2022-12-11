package main

import (
	cmd "emailanlyzer/m/commands"
	"os"

	"github.com/alecthomas/kong"
)

type CLI struct {
	Parse cmd.ParseCLI `cmd:"" help:"parse an eml file"`
}

func main() {

	parser := kong.Must(&CLI{},
		kong.Name("mailp"),
		kong.Description("suspecious emails analyzer"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Tree: true,
		}))

	ctx, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)
	err = ctx.Run()
	ctx.FatalIfErrorf(err)
}
