# agents/face/install_onnx.py

import subprocess, os

try:
    subprocess.run(['nvidia-smi'], check=True, stdout=open(os.devnull, 'w'))
    print("🔧 GPU detected — installing onnxruntime-gpu...")
    subprocess.run(['pip', 'install', '--no-cache-dir', 'onnxruntime-gpu'], check=True)
except (subprocess.CalledProcessError, FileNotFoundError):
    print("⚙️  No GPU or nvidia-smi not found — installing onnxruntime...")
    subprocess.run(['pip', 'install', '--no-cache-dir', 'onnxruntime'], check=True)
