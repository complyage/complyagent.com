#================================================================================================
# Dockerfile :: agent/face (CPU + GPU Compatible)
#================================================================================================

FROM python:3.10-slim

ENV DEBIAN_FRONTEND=noninteractive

# -----------------------------------------------------------------------------------------------
# System dependencies
# -----------------------------------------------------------------------------------------------
RUN apt-get update && apt-get install -y --no-install-recommends \
    ffmpeg \
    libsm6 \
    libxext6 \
    libgl1 \
    libglib2.0-0 \
    ca-certificates \
    build-essential \
    python3-dev \
    g++ \
    wget \
    curl \
    git \
    pciutils \
    && rm -rf /var/lib/apt/lists/*

# -----------------------------------------------------------------------------------------------
# Create working directory
# -----------------------------------------------------------------------------------------------
WORKDIR /app

# -----------------------------------------------------------------------------------------------
# Copy requirements (without onnxruntime)
# -----------------------------------------------------------------------------------------------
COPY requirements.txt ./requirements.txt

# Remove any fixed onnxruntime lines from the file just in case
RUN grep -vE '^onnxruntime(-gpu)?' requirements.txt > requirements.clean.txt

# Install core dependencies (insightface, flask, etc.)
RUN pip install --no-cache-dir --upgrade pip && \
    pip install -r requirements.clean.txt && \
    pip install gunicorn

# -----------------------------------------------------------------------------------------------
# Copy application code
# -----------------------------------------------------------------------------------------------

COPY . .

# -----------------------------------------------------------------------------------------------
# Dynamically install onnxruntime OR onnxruntime-gpu
# -----------------------------------------------------------------------------------------------

RUN python3 install_onnx.py

# -----------------------------------------------------------------------------------------------
# Set PYTHONPATH and start server
# -----------------------------------------------------------------------------------------------
ENV PYTHONPATH=/app

CMD ["gunicorn", "-w", "1", "-b", "0.0.0.0:50052", "server:app"]
