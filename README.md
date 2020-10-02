# cmd_runner

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