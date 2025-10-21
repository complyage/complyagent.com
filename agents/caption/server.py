# server.py :: BLIP Image Caption Agent

from flask import Flask, request
from response import json_error, json_success
from transformers import BlipProcessor, BlipForConditionalGeneration
from PIL import Image
import torch, base64, io, traceback

# Init app + model
app = Flask(__name__)

device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
processor = BlipProcessor.from_pretrained("Salesforce/blip-image-captioning-base")
model     = BlipForConditionalGeneration.from_pretrained("Salesforce/blip-image-captioning-base").to(device)

def decode_image(b64: str) -> Image.Image:
    return Image.open(io.BytesIO(base64.b64decode(b64))).convert("RGB")

@app.route("/caption", methods=["POST"])
def caption_image():
    try:
        data = request.get_json()
        if not data or "image" not in data:
            return json_error("Missing 'image' field")

        img = decode_image(data["image"])
        inputs = processor(images=img, return_tensors="pt").to(device)
        out = model.generate(**inputs, max_new_tokens=50)
        caption = processor.decode(out[0], skip_special_tokens=True)

        return json_success(True, data={"caption": caption})

    except Exception as e:
        return json_error(str(e) + "\n" + traceback.format_exc())

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
