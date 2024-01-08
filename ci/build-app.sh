#!/bin/bash
cd app/restapi
go build -gcflags="-N -l" ./cmd/server/server.go
echo "Server was build successfully."
go build -gcflags="-N -l" ./cmd/client/client.go
echo "Client was build successfully."
go build -gcflags="-N -l" ./cmd/healthchecker/healthchecker.go
echo "Health checker was build successfully."