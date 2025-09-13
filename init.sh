#!/usr/bin/env bash
#==================================================================================================
# init.sh :: Environment Initializer
# Downloads required models and sets up agents
#==================================================================================================

set -e

#--------------------------------------------------------------------------------------------------
# NSFW Model Setup
#--------------------------------------------------------------------------------------------------

MODEL_URL="https://github.com/notAI-tech/NudeNet/releases/download/v3/640m.onnx"
MODEL_DIR="agents/nsfw/model"
MODEL_PATH="$MODEL_DIR/640m.onnx"

echo "||------------------------------------------------------------||"
echo "|| Initializing ComplyAgent Environment                       ||"
echo "||------------------------------------------------------------||"

mkdir -p "$MODEL_DIR"

if [ -f "$MODEL_PATH" ]; then
    echo "✓ NudeNet model already exists at $MODEL_PATH"
else
    echo "↓ Downloading NudeNet model..."
    curl -L "$MODEL_URL" -o "$MODEL_PATH"
    echo "✓ Saved model to $MODEL_PATH"
fi

echo "||------------------------------------------------------------||"
echo "|| Initialization Complete                                    ||"
echo "||------------------------------------------------------------||"
