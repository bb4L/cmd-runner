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

// RunCmd run the command
func RunCmd(command string, args []string) {
	InfoLogger := logging.GetInfoLogger()
	ErrorLogger := logging.GetErrorLogger()
	FatalLogger := logging.GetFatalLogger()

	InfoLogger.Printf("executing: %s", command)
	cmd := exec.Command(command)

	if len(args) > 0 {
		InfoLogger.Printf("using args: %s", args)
		cmd = exec.Command(command, args...)
	}

	cmd.Stdout = InfoLogger.Writer()
	cmd.Stderr = ErrorLogger.Writer()
	err := cmd.Run()
	if err != nil {
		FatalLogger.Fatalln("failed with")
		FatalLogger.Fatalln(err)
	}
}
