#================================================================================================
# Dockerfile :: agent/caption
#================================================================================================

FROM python:3.10-slim

ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    libgl1 \
    libglib2.0-0 \
    git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY requirements.txt ./requirements.txt
RUN pip install --no-cache-dir --upgrade pip && \
    pip install -r requirements.txt && \
    pip install gunicorn

# Download model (optional — forces pre-cache)
COPY install_blip.py ./
RUN python3 install_blip.py

COPY . .

CMD ["gunicorn", "-w", "1", "-b", "0.0.0.0:8080", "server:app"]
