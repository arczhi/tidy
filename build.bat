go env -w GOOS=linux & go build -tags 'linux' -o ./bin/tidy ./cmd
go env -w GOOS=windows & go build -tags 'windows' -o ./bin/tidy.exe ./cmd