#================================================================================================
# route/detect.py :: Face Detection Endpoint
#================================================================================================

from flask            import request
from PIL              import Image
import base64, io, numpy as np, traceback, hashlib

from response         import json_error, json_success

#||------------------------------------------------------------------------------------------------||
#|| Decode Image from Base64
#||------------------------------------------------------------------------------------------------||

def decode_image(b64):
      try:
            img_bytes = base64.b64decode(b64)
            pil_img   = Image.open(io.BytesIO(img_bytes)).convert("RGB")
            return np.array(pil_img)
      except Exception as e:
            raise ValueError(f"Invalid image: {str(e)}")

#||------------------------------------------------------------------------------------------------||
#|| Generate Stable Facial ID (hash of embedding vector)
#||------------------------------------------------------------------------------------------------||

def embedding_to_id(embedding):
      emb_string = ",".join(f"{v:.6f}" for v in embedding)
      return hashlib.sha256(emb_string.encode()).hexdigest()

#||------------------------------------------------------------------------------------------------||
#|| Detect Route Handler
#||------------------------------------------------------------------------------------------------||

def detect_route(face_app):
      def detect_handler():
            try:
                  data    = request.get_json()
                  img_b64 = data.get("image")

                  if not img_b64:
                        return json_error("Missing 'image' field")

                  img    = decode_image(img_b64)
                  faces  = face_app.get(img)

                  if not faces:
                        return json_error("No faces detected", status=200)

                  face_data = []
                  age_list  = []

                  for face in faces:
                        age        = int(face.age)
                        gender     = "male" if face.gender == 1 else "female"
                        confidence = float(face.det_score)
                        embedding  = face.embedding.tolist()
                        face_id    = embedding_to_id(embedding)

                        age_list.append(age)
                        face_data.append({
                              "id"        : face_id,
                              "age"       : age,
                              "gender"    : gender,
                              "confidence": confidence,
                              "embedding" : embedding
                        })

                  data_out = {
                        "count"   : len(face_data),
                        "age_min" : min(age_list),
                        "age_max" : max(age_list),
                        "faces"   : face_data
                  }

                  return json_success(True, "Face(s) detected", data=data_out)

            except Exception as e:
                  tb = traceback.format_exc()
                  return json_error(f"{str(e)}\n{tb}")

      detect_handler.__name__ = "detect_handler"
      return detect_handler
