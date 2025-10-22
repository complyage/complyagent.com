# Base Image
FROM python:3.10-slim

# Set environment variables for non-interactive installs
ENV DEBIAN_FRONTEND=noninteractive

# Install system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    ffmpeg \
    libsm6 \
    libxext6 \
    libgl1 \
    libglib2.0-0 \
    ca-certificates \
    ccache \
    wget \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory to /app
WORKDIR /app

# Copy requirements and install Python dependencies
COPY requirements.txt ./requirements.txt
RUN pip install --no-cache-dir --upgrade pip && \
    pip install -r requirements.txt

# Copy application source code
COPY . .

# Set PYTHONPATH to /app
ENV PYTHONPATH=/app

# Default command to run the server
CMD ["gunicorn", "-w", "1", "-b", "0.0.0.0:8080", "server:app"]


