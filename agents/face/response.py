#||------------------------------------------------------------------------------------------------||
#|| Response Formatter
#|| Action to compare faces (with error handling)
#||------------------------------------------------------------------------------------------------||

from flask import jsonify
from time import time
from typing import Optional, Any

#||------------------------------------------------------------------------------------------------||
#|| Detection Successful
#||------------------------------------------------------------------------------------------------||

def detect_success(
      age: int,
      gender: str,
      confidence: Optional[float] = None,
):
      data = {
            "age"       : age,
            "gender"    : gender,
            "confidence": confidence
      }

      return json_success(True, "Face detected", data=data)
  
#||------------------------------------------------------------------------------------------------||
#|| Match was successful
#||------------------------------------------------------------------------------------------------||
  
def compare_success(
      match: bool,
      confidence: Optional[float] = None,
):
      data = {
            "match"       : match,
            "confidence": confidence
      }

      return json_success(True, "", data=data)  

#||------------------------------------------------------------------------------------------------||
#|| Base Error
#||------------------------------------------------------------------------------------------------||

def json_error(message: str):
      response = {
            "success": False,
            "message": message,
            "data": {}
      }

      return jsonify(response), 400

#||------------------------------------------------------------------------------------------------||
#|| Base JSON Response
#||------------------------------------------------------------------------------------------------||

def json_success(
      success: True,
      message: Optional[str] = None,
      data: Optional[Any] = None,
):
      response = {
            "success": success,
            "message": message or ("OK" if success else "Error"),
            "data": data or {}
      }

      return jsonify(response), 200

#||------------------------------------------------------------------------------------------------||
#|| Import
#||------------------------------------------------------------------------------------------------||


