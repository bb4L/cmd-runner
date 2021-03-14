# cmd_runner
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/bb4L/cmd_runner)
[![CI](https://github.com/bb4L/cmd_runner/actions/workflows/build.yml/badge.svg)](https://github.com/bb4L/cmd_runner/actions/workflows/build.yml)
![GitHub last commit](https://img.shields.io/github/last-commit/bb4L/cmd_runner)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/bb4L/cmd_runner)
![GitHub Release Date](https://img.shields.io/github/release-date/bb4L/cmd_runner)
![GitHub](https://img.shields.io/github/license/bb4L/cmd_runner)

## Build
run the build script `build.sh`

## Configuration
A configuration `yaml`-file should be passed as the first argument to the programm.

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

**For windows users** pay attention that the fullpath of the programm should escape the `\`

# License
[GPLv3](LICENSE)
