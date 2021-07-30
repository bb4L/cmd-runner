package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bb4L/cmd_runner/executor"
	"github.com/bb4L/cmd_runner/logging"

	"gopkg.in/yaml.v2"
)

var logger = logging.GetLogger(os.Stdout, "cmd_runner", "runner")

type config struct {
	Cmds []executor.Command `yaml:"cmds"`
}

func main() {
	configPath := flag.String("c", "./config.yaml", "path to config file")
	flag.Parse()

	logger.Printf("starting with config in: %s", *configPath)

	data, err := ioutil.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	var cmds config

	source := []byte(data)
	err = yaml.Unmarshal(source, &cmds)
	if err != nil {
		logger.Fatalf("error: %v", err)
	}

	executor.RunCmds(cmds.Cmds)

	buf := bufio.NewReader(os.Stdin)
	logger.Println("press q to close")

	for closingData, err := buf.ReadString('\n'); strings.TrimRight(strings.TrimRight(closingData, "\r\n"), "\n") != "q" || err != nil; closingData, err = buf.ReadString('\n') {
		if err != nil {
			logger.Println(err)
			continue
		}
	}
}
