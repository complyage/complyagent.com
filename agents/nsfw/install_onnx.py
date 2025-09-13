#================================================================================================
# install_onnx.py :: Auto-install ONNXRuntime (GPU if available, else CPU)
#================================================================================================

import subprocess
import os

try:
    # Check if nvidia-smi exists and works
    subprocess.run(["nvidia-smi"], check=True, stdout=open(os.devnull, "w"))
    print("GPU detected — installing onnxruntime-gpu")
    subprocess.run(["pip", "install", "--no-cache-dir", "onnxruntime-gpu"], check=True)
except Exception:
    print("No GPU detected — installing onnxruntime (CPU)")
    subprocess.run(["pip", "install", "--no-cache-dir", "onnxruntime"], check=True)
