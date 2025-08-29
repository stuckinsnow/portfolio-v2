#!/bin/bash

echo "Building portfolio..."

# Install dependencies
pnpm install

# Build frontend (creates dist/ directory)
pnpm run build

echo "Build complete!"
echo "Now run: ./run.sh"