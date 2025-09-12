#================================================================================================
# route/compare.py :: Face Comparison Endpoint
#================================================================================================

from flask            import request
from PIL              import Image
import base64, io, numpy as np, traceback

from response         import json_error, match_success

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
#|| Compare Route Handler
#||------------------------------------------------------------------------------------------------||

def compare_route(face_app):
      def compare_handler():
            try:
                  data   = request.get_json()
                  b64_1  = data.get("image1")
                  b64_2  = data.get("image2")

                  if not b64_1 or not b64_2:
                        return json_error("Missing 'image1' or 'image2' field")

                  img1   = decode_image(b64_1)
                  img2   = decode_image(b64_2)

                  faces1 = face_app.get(img1)
                  faces2 = face_app.get(img2)

                  if not faces1 or not faces2:
                        return match_success(False)

                  emb1   = faces1[0].embedding
                  emb2   = faces2[0].embedding

                  score  = np.dot(emb1, emb2) / (np.linalg.norm(emb1) * np.linalg.norm(emb2))
                  match  = score > 0.35

                  return match_success(match, confidence=score)

            except Exception as e:
                  tb = traceback.format_exc()
                  return json_error(f"{str(e)}\n{tb}")

      compare_handler.__name__ = "compare_handler"
      return compare_handler
