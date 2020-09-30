package main

import (
	"cmdrunner/executor"
	"cmdrunner/logging"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	err error
)

type config struct {
	Cmds []executor.Command `yaml:"cmds"`
}

func init() {
}
func main() {
	InfoLogger := logging.GetInfoLogger()
	FatalLogger := logging.GetFatalLogger()

	argsWithoutProg := os.Args[1:]

	InfoLogger.Printf("starting with %s", argsWithoutProg)

	if len(argsWithoutProg) != 1 {
		FatalLogger.Fatalln("wrong number of arguments, exactly 1 is needed")
	}
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var cmds config

	source := []byte(data)
	err = yaml.Unmarshal(source, &cmds)
	if err != nil {
		FatalLogger.Fatalf("error: %v", err)
	}

	for _, cmd := range cmds.Cmds {
		executor.RunCmd(cmd.Cmd, cmd.Args)
	}
}
