#================================================================================================
# server.py :: NSFW Detection Agent (Flask + NudeNet v3.x using NudeDetector)
#================================================================================================

from flask import Flask, request
from response import json_error, nsfw_success
from nudenet import NudeDetector
import base64, io, traceback, os, pathlib

#------------------------------------------------------------------------------------------------
# Initialize App + Model
#------------------------------------------------------------------------------------------------

app = Flask(__name__)

BASE_DIR   = pathlib.Path(__file__).resolve().parent
MODEL_PATH = str((BASE_DIR / "model/640m.onnx").resolve())  # ✅ model in /app root

print("||-------- NSFW Agent Paths ---------------------||")
print(f"|| BASE_DIR   : {BASE_DIR}")
print(f"|| MODEL_PATH : {MODEL_PATH}")
print(f"|| CWD        : {os.getcwd()}")
print(f"|| Model exists? {os.path.exists(MODEL_PATH)}")
print("||------------------------------------------------||")

detector = NudeDetector(model_path=MODEL_PATH, inference_resolution=640)

#------------------------------------------------------------------------------------------------
# Decode Image Helper
#------------------------------------------------------------------------------------------------

def decode_image(b64: str) -> str:
    img_bytes = base64.b64decode(b64)
    tmp_path  = "/tmp/upload.jpg"
    with open(tmp_path, "wb") as f:
        f.write(img_bytes)
    return tmp_path

#------------------------------------------------------------------------------------------------
# Route: /detect
#------------------------------------------------------------------------------------------------

@app.route("/detect", methods=["POST"])
def detect_nsfw():
    try:
        data = request.get_json()
        if not data:
            return json_error("Invalid JSON")

        b64 = data.get("image")
        if not b64:
            return json_error("Missing 'image' field")

        img_path = decode_image(b64)
        detections = detector.detect(img_path)

        print(f"[NSFWAgent] Detections: {detections}")  # ✅ debug log

        if not detections or not isinstance(detections, list):
            return nsfw_success(nsfw=False, confidence=0.0, classifications={})

        classifications = {}
        max_conf = 0.0
        nsfw = False

        SFW_LABELS = [
            "ANUS_COVERED",
            "ARMPITS_COVERED",
            "ARMPITS_EXPOSED",
            "BELLY_COVERED",
            "BELLY_EXPOSED",
            "BUTTOCKS_COVERED",
            "FACE_FEMALE",
            "FACE_MALE",
            "FEMALE_BREAST_COVERED",
            "FEET_COVERED",
            "FEET_EXPOSED",
            "MALE_BREAST_EXPOSED",
            "MALE_GENITALIA_COVERED"
        ]

        for det in detections:
            label = det.get("class") or det.get("label")
            score = float(det.get("score", 0.0))
            if not label:
                continue

            classifications[label] = max(classifications.get(label, 0.0), score)

            if label not in SFW_LABELS:  # ✅ Only flip NSFW for true unsafe detections
                nsfw = True

            max_conf = max(max_conf, score)

        return nsfw_success(nsfw, max_conf, classifications)

    except Exception as e:
        tb = traceback.format_exc()
        return json_error(f"{str(e)}\n{tb}")

#------------------------------------------------------------------------------------------------
# Entrypoint
#------------------------------------------------------------------------------------------------

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
