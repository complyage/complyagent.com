#!/bin/sh
set -e

echo ">>> Starting Ollama server..."
ollama serve &

# Wait a few seconds for the server to start
sleep 5

# Pull phi3 model if not cached
if ! ollama list | grep -q "phi3"; then
  echo ">>> Pulling phi3 model..."
  ollama pull phi3
else
  echo ">>> phi3 already available"
fi

# Keep server running
wait
