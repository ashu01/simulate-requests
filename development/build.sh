#!/bin/bash
begin=$(date +%s)
echo "Building for mac"
env GOOS=darwin GOARCH=amd64 go build -v -o connection-mac connection.go &


echo "Building for linux"
env GOOS=linux GOARCH=amd64 go build -v -o connection-linux connection.go &


echo "Building for windows"
env GOOS=windows GOARCH=amd64 go build -v -o connection-win.exe connection.go &

wait

end=$(date +%s)

total=$((end-begin))

echo "Built-in "$total"s"
