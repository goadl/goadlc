package main

import (
	"fmt"
	"os"

	"github.com/goadl/goadlc/pkg/goadlc"
	"github.com/goadl/goadlc/pkg/types"
	"github.com/jpillora/opts"
)

// Set by build tool chain by
// go build --ldflags '-X main.XXX=xxx -X main.YYY=yyy -X main.ZZZ=zzz'
var (
	Version    string = "dev"
	Date       string = "na"
	Commit     string = "na"
	ReleaseURL string = "na"
)

type versionCmd struct{}

func (r *versionCmd) Run() {
	fmt.Printf(`
version: %s
date:    %s
commit:  %s
release: %s
`, Version, Date, Commit, ReleaseURL)
}

var (
	rflg    = &types.Root{}
	cliBldr = opts.New(rflg).
		Name("goadlc").
		EmbedGlobalFlagSet().
		Complete()
)

func main() {
	cli, err := cliBldr.ParseArgsError(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\nError: %v\n\n", cli.Selected().Help(), err)
		os.Exit(1)
	}
	err = cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\nError: %v\n\n", cli.Selected().Help(), err)
		os.Exit(1)
	}
}

func init() {
	cliBldr.AddCommand(opts.New(&versionCmd{}).Name("version"))
}

func init() {
	cliBldr.
		AddCommand(opts.New(goadlc.NewAst(rflg)).Name("ast").
			End())
}
