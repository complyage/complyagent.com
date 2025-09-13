#================================================================================================
# response.py :: Standardized JSON Responses
#================================================================================================

from flask import jsonify
from typing import Optional, Any

#||------------------------------------------------------------------------------------------------||
#|| NSFW Detection Success
#||------------------------------------------------------------------------------------------------||

def nsfw_success(nsfw: bool, confidence: float, classifications: Optional[dict] = None):
    return json_success(True, "", data={
        "nsfw": nsfw,
        "confidence": float(confidence),
        "classifications": classifications or {}
    })

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
#|| Base JSON Success
#||------------------------------------------------------------------------------------------------||

def json_success(success: bool, message: Optional[str] = None, data: Optional[Any] = None):
    response = {
        "success": success,
        "message": message or ("OK" if success else "Error"),
        "data": data or {}
    }
    return jsonify(response), 200
