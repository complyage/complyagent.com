#================================================================================================
# OCR HTTP Server (Flask)
#================================================================================================

from flask import Flask, request, jsonify
from paddleocr import PaddleOCR
from PIL import Image
import base64
import io
import paddle
import numpy as np
import traceback

#------------------------------------------------------------------------------------------------
# Initialize Flask app and OCR model
#------------------------------------------------------------------------------------------------

app = Flask(__name__)

use_gpu = paddle.device.is_compiled_with_cuda()
device = "gpu" if use_gpu else "cpu"

ocr_model = PaddleOCR(
    lang="en",
    device=device,
    ocr_version="PP-OCRv5",
    use_textline_orientation=True,
    use_doc_orientation_classify=False,
    use_doc_unwarping=False
)

#------------------------------------------------------------------------------------------------
# Route /ocr
#------------------------------------------------------------------------------------------------

#------------------------------------------------------------------------------------------------
# Route /ocr
#------------------------------------------------------------------------------------------------

from response import json_success, json_error, ocr_success

@app.route("/ocr", methods=["POST"])
def extract_text():
      try:
            data = request.get_json()
            if not data:
                  return json_error("Invalid JSON")

            b64 = data.get("image")
            if not b64:
                  return json_error("Missing 'image' field")

            try:
                  img_bytes = base64.b64decode(b64)
                  pil_img   = Image.open(io.BytesIO(img_bytes))
            except Exception as e:
                  return json_error(f"Invalid image: {str(e)}")

            if pil_img.mode != "RGB":
                  pil_img = pil_img.convert("RGB")

            img_np = np.array(pil_img, dtype=np.uint8)

            result     = ocr_model.predict(img_np)
            texts      = result[0].get("rec_texts", [])
            final_text = " ".join(texts) if texts else ""

            return ocr_success(final_text)

      except Exception as e:
            tb = traceback.format_exc()
            return json_error(f"{str(e)}\n{tb}")