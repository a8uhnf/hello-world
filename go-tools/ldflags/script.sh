go build -ldflags="-X 'main.Version=v1.0.0'" # change variable name of main package's Version variable
go tool nm ./ldflags | grep app # find variables which has path prefix app