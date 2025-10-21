# response.py :: Shared JSON response helpers

from flask import jsonify
from typing import Optional, Any

def json_error(message: str):
    return jsonify({"success": False, "message": message, "data": {}}), 400

def json_success(success: bool, message: Optional[str] = None, data: Optional[Any] = None):
    return jsonify({
        "success": success,
        "message": message or ("OK" if success else "Error"),
        "data": data or {}
    }), 200
