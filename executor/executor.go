package executor

import (
	"os"
	"os/exec"

	"github.com/bb4L/cmd-runner/logging"
)

var logger = logging.GetLogger(os.Stdout, "cmd-runner", "executor")

// Command struct for the command to be executed
type Command struct {
	Cmd  string   `yaml:"cmd"`
	Args []string `yaml:"args"`
}

// RunCmds run all the commands
func RunCmds(cmds []Command) {
	for _, cmd := range cmds {
		logger.Printf("starting %s with %s", cmd.Cmd, cmd.Args)
		RunCmd(cmd.Cmd, cmd.Args)
	}
}

// RunCmd run the command
func RunCmd(command string, args []string) {
	var cmd *exec.Cmd

	if len(args) > 0 {
		cmd = exec.Command(command, args...)
	} else {
		cmd = exec.Command(command)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()

	if err != nil {
		logger.Fatalf("failed with: %s", err)
	}
}
