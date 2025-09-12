#!/bin/bash
set -e

# If mounted volume is empty, copy pre-downloaded models into it
if [ -z "$(ls -A /root/.paddleocr 2>/dev/null)" ]; then
    echo "📥 Populating PaddleOCR models into /root/.paddleocr volume..."
    mkdir -p /root/.paddleocr
    cp -r /app/.paddleocr/* /root/.paddleocr/
else
    echo "✅ PaddleOCR models already present in volume, skipping copy."
fi

# Hand off to CMD
exec "$@"
