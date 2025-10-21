#================================================================================================
# serve.py :: Agent Face Server Entry Point
#================================================================================================

from flask import Flask
from insightface.app import FaceAnalysis

#||------------------------------------------------------------------------------------------------||
#|| App Init
#||------------------------------------------------------------------------------------------------||

app = Flask(__name__)

face_app = FaceAnalysis(providers=[
      'CUDAExecutionProvider',     # Try GPU
      'CPUExecutionProvider'       # Fallback to CPU
])
face_app.prepare(ctx_id=0)

#||------------------------------------------------------------------------------------------------||
#|| Routes
#||------------------------------------------------------------------------------------------------||

from route_detect import detect_route
from route_compare import compare_route

app.add_url_rule("/detect",  view_func=detect_route(face_app),  methods=["POST"])
app.add_url_rule("/compare", view_func=compare_route(face_app), methods=["POST"])

#||------------------------------------------------------------------------------------------------||
#|| Entry Point
#||------------------------------------------------------------------------------------------------||

if __name__ == "__main__":
      app.run(host="0.0.0.0", port=8080)
