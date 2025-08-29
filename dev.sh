#!/bin/bash

echo "Starting development server with auto-rebuild..."

# Check if sass is installed
if ! command -v sass &> /dev/null; then
    echo "Installing sass..."
    if command -v npm &> /dev/null; then
        npm install -g sass
    else
        echo "Error: npm not found. Please install Node.js and npm first."
        echo "Or install sass manually with your package manager."
        exit 1
    fi
fi

# Create CSS file initially
echo "Initial SCSS compilation..."
cd static/css
sass style.scss style.css
cd ../..

# Start sass watcher in background
echo "Starting SCSS watcher..."
cd static/css
sass --watch style.scss:style.css &
SASS_PID=$!
cd ../..

# Function to cleanup on exit
cleanup() {
    echo "Stopping development server..."
    kill $SASS_PID 2>/dev/null
    exit 0
}

trap cleanup SIGINT SIGTERM

# Start Go server
cd backend
echo "Server starting at http://localhost:8080"
echo "Press Ctrl+C to stop"
go run .