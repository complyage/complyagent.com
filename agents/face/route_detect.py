#================================================================================================
# route/detect.py :: Face Detection Endpoint
#================================================================================================

from flask            import request
from PIL              import Image
import base64, io, numpy as np, traceback

from response         import json_error, detect_success

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
#|| Detect Route Handler
#||------------------------------------------------------------------------------------------------||

def detect_route(face_app):
      def detect_handler():
            try:
                  data     = request.get_json()
                  img_b64  = data.get("image")

                  if not img_b64:
                        return json_error("Missing 'image' field")

                  img     = decode_image(img_b64)
                  faces   = face_app.get(img)

                  if not faces:
                        return json_error("No face detected", status=200)

                  face    = faces[0]

                  return detect_success(
                        age       = int(face.age),
                        gender    = "male" if face.gender == 1 else "female",
                        confidence= float(face.det_score)
                  )

            except Exception as e:
                  tb = traceback.format_exc()
                  return json_error(f"{str(e)}\n{tb}")

      detect_handler.__name__ = "detect_handler"
      return detect_handler
