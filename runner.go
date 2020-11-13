package main

import (
	"bufio"
	"cmdrunner/executor"
	"cmdrunner/logging"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	err error
)

type config struct {
	Cmds []executor.Command `yaml:"cmds"`
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

	executor.RunCmds(cmds.Cmds)

	buf := bufio.NewReader(os.Stdin)
	InfoLogger.Print("press q to close")

	for closingData, err := buf.ReadString('\n'); strings.TrimRight(strings.TrimRight(closingData, "\r\n"), "\n") != "q" || err != nil; closingData, err = buf.ReadString('\n') {
		if err != nil {
			FatalLogger.Println(err)
			continue
		}
	}
}
