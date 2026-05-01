#!/bin/sh

echo "Running pre-start commands..."


echo "Starting main app..."
exec ./go-api-app

# your custom commands
cd /workspace
go run .   # or any setup task