package executor

import (
	"cmdrunner/logging"
	"os/exec"
)

// Command struct for the command to be executed
type Command struct {
	Cmd  string   `yaml:"cmd"`
	Args []string `yaml:"args"`
}

// RunCmds run all the commands
func RunCmds(cmds []Command) {
	InfoLogger := logging.GetInfoLogger()

	for _, cmd := range cmds {
		InfoLogger.Printf("starting %s with %s", cmd.Cmd, cmd.Args)
		RunCmd(cmd.Cmd, cmd.Args)
	}
}

// RunCmd run the command
func RunCmd(command string, args []string) {
	InfoLogger := logging.GetInfoLogger()
	ErrorLogger := logging.GetErrorLogger()
	FatalLogger := logging.GetFatalLogger()

	InfoLogger.Printf("executing: %s", command)
	var cmd *exec.Cmd

	if len(args) > 0 {
		InfoLogger.Printf("using args: %s", args)
		cmd = exec.Command(command, args...)
	} else {
		cmd = exec.Command(command)
	}

	cmd.Stdout = InfoLogger.Writer()
	cmd.Stderr = ErrorLogger.Writer()
	err := cmd.Start()

	if err != nil {
		FatalLogger.Printf("failed with: %s", err)
	}
}
