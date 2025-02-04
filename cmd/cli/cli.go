package main

import (
	// "flag"
	"fmt"
	"os"
)

type Cmd int

const (
	CmdTimeCapsule Cmd = iota
	CmdMetrics
	CmdSearch
	CmdWebServer
)

func (cmd Cmd) handle() {
	switch cmd {
	case CmdTimeCapsule:
		panic("TODO")
		break
	case CmdMetrics:
		panic("TODO")
		break
	case CmdSearch:
		panic("TODO")
		break
	case CmdWebServer:
		panic("TODO")
		break
	default:
		panic("Unreachable")
	}
}

// aliases for the Cmd
var Cmds = map[string]Cmd{
	"time-capsule": CmdTimeCapsule,
	"capsule":      CmdTimeCapsule,
	"t":            CmdTimeCapsule,
	"tc":           CmdTimeCapsule,

	"metrics": CmdMetrics,
	"m":       CmdMetrics,

	"search": CmdSearch,
	"s":      CmdSearch,

	"web-server": CmdWebServer,
	"w":          CmdWebServer,
	"ws":         CmdWebServer,
}

type Cli struct {
	program string
	command Cmd
}

func main() {
	cli, err := parseArgs(os.Args)

	if err != nil {
		fmt.Printf("%v\n%s", err, usage(os.Args[0]))
		os.Exit(1)
	}

	cli.command.handle()
}

func usage(program string) string {
	return fmt.Sprintf("USAGE: %s <command> [<args>]\n", program)
}

func parseArgs(args []string) (*Cli, error) {
	program, programArgs := args[0], args[1:]

	if len(programArgs) < 1 {
		return nil, fmt.Errorf("ERROR: Missing command\n")
	}

	cmd, ok := Cmds[programArgs[0]]

	if !ok {
		return nil, fmt.Errorf("ERROR: Unknown command: %s\n", programArgs[0])
	}

	return &Cli{
		program: program,
		command: cmd,
	}, nil
}
