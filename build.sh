GOOS=windows GOARCH=amd64 go build -o runner_win64.exe runner.go 
GOOS=linux GOARCH=amd64 go build -o runner_linux64 runner.go 