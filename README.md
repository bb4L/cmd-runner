# cmd_runner
[![Go Reference](https://pkg.go.dev/badge/github.com/bb4L/rpi-radio-alarm-go-library.svg)](https://pkg.go.dev/github.com/bb4L/cmd_runner)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/bb4L/cmd_runner)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/bb4L/cmd_runner)
![GitHub Release Date](https://img.shields.io/github/release-date/bb4L/cmd_runner)
![GitHub last commit](https://img.shields.io/github/last-commit/bb4L/cmd_runner)
![GitHub](https://img.shields.io/github/license/bb4L/cmd_runner)
[![CI](https://github.com/bb4L/cmd_runner/actions/workflows/build.yml/badge.svg)](https://github.com/bb4L/cmd_runner/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bb4L/cmd_runner)](https://goreportcard.com/report/github.com/bb4L/cmd_runner)

## Build
run the build script `build.sh`

## Configuration
A configuration `yaml`-file should be passed as the first argument with the flag `-c` to the programm.

It should have the following structure:

```yaml
cmds:
    - 
      cmd: "programm_to_run"
      args:
        - "arg1"
        - "arg2"

    ...
```

# License
[GPLv3](LICENSE)
