echo installing task
go install github.com/go-task/task/v3/cmd/task@latest
echo installing air
go install github.com/cosmtrek/air@latest
export PORT=8080