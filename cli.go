package main

import (
	"fmt"
	"os"

	"github.com/semiherdogan/aws-ssh/internal/pkg/utils"
)

type fn[T any] func()

type Cli struct {
	Args    []string
	Version string
}

func NewCli() *Cli {
	return &Cli{
		Args:    os.Args[1:],
		Version: "0.2.2",
	}
}

func (cli *Cli) Run() error {
	if utils.Contains(cli.Args, "-h") || utils.Contains(cli.Args, "--help") {
		cli.ShowHelp()
		os.Exit(0)
	}

	if utils.Contains(cli.Args, "-v") || utils.Contains(cli.Args, "--version") {
		cli.ShowVersion()
		os.Exit(0)
	}

	return nil
}

func (cli *Cli) AddCommand(long string, short string, command fn[interface{}]) {
	if utils.Contains(cli.Args, long) || utils.Contains(cli.Args, short) {
		command()
	}

	cli.RemoveArg(long)
	cli.RemoveArg(short)
}

func (cli *Cli) ShowHelp() {
	helpString := `USAGE:
   aws-ssh [--profile] [--region] argument1 argument2 ...

VERSION:
   %s

DESCRIPTION:
   Easy way to connect aws ec2 through ssm
   For more info: https://github.com/semiherdogan/aws-ssh

OPTIONS:
   --profile, -p  show profile select (default: false)
   --region, -r   show region select (default: false)
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
`

	fmt.Printf(helpString, cli.Version)
}

func (cli *Cli) ShowVersion() {
	fmt.Printf("aws-ssh %s\n", cli.Version)
}

func (cli *Cli) RemoveArg(arg string) {
	cli.Args = utils.Remove(cli.Args, arg)
}
